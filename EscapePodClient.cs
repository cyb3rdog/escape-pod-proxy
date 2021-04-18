using Anki.Vector.GrpcUtil;
using Cybervector;
using Grpc.Core;
using System;
using System.Threading;
using System.Threading.Tasks;
using static Cybervector.CyberVectorProxyService;

namespace Cyb3rPod
{
    /// <summary>
    /// Proxy event args
    /// </summary>
    /// <seealso cref="System.EventArgs" />
    public class ProxyMessageEventArgs : EventArgs
    {
        /// <summary>
        /// Gets the event type.
        /// </summary>
        public ProxyMessaage ProxyMessaage { get; }

        /// <summary>
        /// Initializes a new instance of the <see cref="ProxyMessageEventArgs"/> class.
        /// </summary>
        /// <param name="eventType">Type of the event.</param>
        public ProxyMessageEventArgs(ProxyMessaage message)
        {
            this.ProxyMessaage = message;
        }
    }

    public class ProxyExceptionEventArgs : EventArgs
    {
        public Exception ProxyException { get; }

        public ProxyExceptionEventArgs(Exception exception)
        {
            this.ProxyException = exception;
        }
    }

    public class EscapePodClient : IDisposable
    {
        /// <summary>
        /// The event feed loop
        /// </summary>
        private IAsyncEventLoop eventFeed;

        internal const int DefaultConnectionTimeout = 5_000;

        private Channel channel;
        private CyberVectorProxyServiceClient client;

        /// <summary>
        /// Gets a value indicating whether this instance is processing events.
        /// </summary>
        public bool IsProccessingEvents => eventFeed.IsActive;

        /// <summary>
        /// Gets the cancellation token source for timing out the event loop
        /// </summary>
        private CancellationTokenSource timeoutCancellationTokenSource = null;

        internal static DateTime GrpcDeadline(int timeout) => DateTime.UtcNow.AddMilliseconds(timeout);

        public event EventHandler<ProxyExceptionEventArgs> OnProxyException;
        public event EventHandler<ProxyMessageEventArgs> OnProxyMessageReceived;

        public event EventHandler Subscribed;
        public event EventHandler Unsubscribed;
        public event EventHandler<string> KeepAlive;
        public event EventHandler<string> IntentHeard;


        public EscapePodClient()
        {

        }

        /// <summary>
        /// Creates the GRPC channel.
        /// </summary>
        /// <param name="hostAndPort">The host and port.</param>
        /// <param name="timeout">The timeout.</param>
        /// <returns>A task that represents the asynchronous operation; the task result contains the connected channel.</returns>
        /// <exception cref="VectorNotFoundException">Unable to establish a connection to Vector.</exception>
        private async Task<Channel> ConnectChannel(string hostAndPort, int timeout)
        {
            // Append the port if it doesn't exist
            if (!hostAndPort.Contains(":")) hostAndPort += ":8090";

            // Create the channel
            //Channel channel = new Channel(hostAndPort, channelCredentials,
            //    new ChannelOption[] { new ChannelOption("grpc.ssl_target_name_override", robotName) }
            //);
            Channel channel = new Channel(hostAndPort, ChannelCredentials.Insecure);

            // Connect to the channel
            try
            {
                // Open the channel
                await channel.ConnectAsync(GrpcDeadline(timeout)).ConfigureAwait(false);
            }
            catch (Exception ex)
            {
                // If failed to open channel throw exception
                channel?.ShutdownAsync();
                throw new Exception("Unable to establish a connection to EscapePod.", ex);
            }
            return channel;
        }

        public async Task Connect(string target)
        {
            this.channel = await this.ConnectChannel(target, DefaultConnectionTimeout);
            this.client = new CyberVectorProxyServiceClient(channel);
            this.timeoutCancellationTokenSource = new CancellationTokenSource();

            this.eventFeed = new AsyncEventLoop<ProxyMessaage>(
                (token) => client.Subscribe(new SubscribeRequest() { KeepAlive = 60 },
                    cancellationToken: timeoutCancellationTokenSource.Token),
                this.ProcessMessage,
                this.ProcessInterrupted,
                this.ProcessException
            );
            await eventFeed.Start().ConfigureAwait(false);
        }

        public void Disconnect()
        {
            if ((client == null) || (channel == null)) return;
            this.client.UnSubscribe(new UnsubscribeRequest());
        }

        private void ProcessMessage(ProxyMessaage proxyMessage)
        {
            Console.WriteLine(proxyMessage);

            this.OnProxyMessageReceived?.Invoke(this, new ProxyMessageEventArgs(proxyMessage));

            switch (proxyMessage.MessageType)
            {
                case MessageType.Subscribed:
                    this.Subscribed?.Invoke(this, new EventArgs()); break;
                case MessageType.Unsubscribed:
                    this.Unsubscribed?.Invoke(this, new EventArgs()); break;
                case MessageType.KeepAlive:
                    this.KeepAlive?.Invoke(this, proxyMessage.MessageData); break;
                case MessageType.ProcessIntent:
                    this.IntentHeard?.Invoke(this, proxyMessage.IntentName); break;
            }
        }

        private void ProcessException(Exception ex)
        {
            this.OnProxyException?.Invoke(this, new ProxyExceptionEventArgs(ex));
        }

        private void ProcessInterrupted()
        {
            throw new Exception("Connection Lost!");
        }

        public async Task<SelectIntentResponse> SelectIntents(string filter_json = "")
        {
            SelectIntentResponse response = await this.RunMethod(client => client.SelectIntentsAsync(new SelectIntentRequest()
            {
                FilterJson = filter_json
            })).ConfigureAwait(false);
            if (response.Response.Code == ResponseMessage.Types.ResponseCode.Failure)
                new Exception(response.Response.Message);
            return response;
        }

        public async Task<InsertIntentResponse> InsertIntent(string intentJson)
        {
            InsertIntentResponse response = await this.RunMethod(client => client.InsertIntentAsync(new InsertIntentRequest()
            {
                IntentData = intentJson
            })).ConfigureAwait(false);
            if (response.Response.Code == ResponseMessage.Types.ResponseCode.Failure)
                new Exception(response.Response.Message);
            return response;
        }

        public async Task<DeleteIntentResponse> DeleteIntent(string intentId)
        {
            DeleteIntentResponse response = await this.RunMethod(client => client.DeleteIntentAsync(new DeleteIntentRequest()
            {
                IntentId = intentId
            })).ConfigureAwait(false);
            if (response.Response.Code == ResponseMessage.Types.ResponseCode.Failure)
                new Exception(response.Response.Message);
            return response;
        }

        /// <summary>
        /// Runs the client method.
        /// </summary>
        /// <typeparam name="T">The result type of the command</typeparam>
        /// <param name="command">The command.</param>
        /// <returns>A task that represents the asynchronous operation; the task result contains the result of the command.</returns>
        internal Task<T> RunMethod<T>(Func<CyberVectorProxyServiceClient, AsyncUnaryCall<T>> command)
        {
            return RunMethod(robot => command(robot).ResponseAsync);
        }

        /// <summary>
        /// Runs the client method.
        /// </summary>
        /// <typeparam name="T">The result type of the command</typeparam>
        /// <param name="command">The command.</param>
        /// <returns>A task that represents the asynchronous operation; the task result contains the result of the command.</returns>
        internal async Task<T> RunMethod<T>(Func<CyberVectorProxyServiceClient, Task<T>> command)
        {
            // Cannot run method if not connected
            if (client == null) throw new Exception("Not connected to EscapePod.");
            return await this.RunCommand(command).ConfigureAwait(false);
        }

        /// <summary>
        /// Runs a command against the gRPC client instance.
        /// </summary>
        /// <typeparam name="TResult">The type of the result.</typeparam>
        /// <param name="command">The command.</param>
        /// <returns>A task that represents the asynchronous operation; the task result contains the result of the command.</returns>
        public async Task<TResult> RunCommand<TResult>(Func<CyberVectorProxyServiceClient, Task<TResult>> command)
        {
            try
            {
                return await command(client).ConfigureAwait(false);
            }
            catch (Grpc.Core.RpcException ex)
            {
                throw ex;
                //throw TranslateGrpcException(ex);
            }
        }

        public void Dispose()
        {
            this.Disconnect();
            client = null;
            channel = null;
        }
    }
}