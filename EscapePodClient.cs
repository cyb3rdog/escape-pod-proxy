using Anki.Vector.GrpcUtil;
using Cybervector;
using Grpc.Core;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Net.Sockets;
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
        private bool isDisconnecting;

        /// <summary>
        /// The event feed loop
        /// </summary>
        private IAsyncEventLoop eventFeed;

        /// <summary>
        /// Gets the cancellation token source for timing out the event loop
        /// </summary>
        private CancellationTokenSource cancellationTokenSource = null;

        private const int DefaultConnectionTimeout = 5_000;
        private static DateTime GrpcDeadline(int timeout) => DateTime.UtcNow.AddMilliseconds(timeout);

        private string uuid;
        private Channel channel;
        private CyberVectorProxyServiceClient client;

        /// <summary>
        /// Gets a value indicating whether this instance is processing events.
        /// </summary>
        public bool IsConnected => (eventFeed != null) && (eventFeed.IsActive);

        public event EventHandler<ProxyExceptionEventArgs> OnProxyException;
        public event EventHandler<ProxyMessageEventArgs> OnProxyMessageReceived;

        public event EventHandler Connected;
        public event EventHandler Disconnected;

        public event EventHandler<SubscribedMessage> Subscribed;
        public event EventHandler Unsubscribed;
        public event EventHandler<KeepAliveMessage> KeepAlive;
        public event EventHandler<string> IntentHeard;

        public string HostName { get; set; }


        public EscapePodClient(string hostName = null)
        {
            this.HostName = hostName;
        }
        ~EscapePodClient()
        {
            this.Dispose();
        }


        public static IPAddress GetIPsByName(string hostName)
        {
            string[] hostAndPort = hostName.Split(new string[] { ":" }, StringSplitOptions.RemoveEmptyEntries);
            if (hostAndPort.Length == 0) return null;

            IPAddress outIpAddress;
            if (IPAddress.TryParse(hostAndPort[0], out outIpAddress) == true)
                return outIpAddress;

#if (NETSTANDARD1_6 || NETCOREAPP1_1)
            IPHostEntry Host = await Dns.GetHostEntryAsync(hostAndPort[0]);
#else
            IPHostEntry Host = Dns.GetHostEntry(hostAndPort[0]);
#endif

            IPAddress[] addresslist = Dns.GetHostAddresses(hostName);

            if (addresslist == null || addresslist.Length == 0)
                return null;

            return addresslist.Where(o => o.AddressFamily == AddressFamily.InterNetwork).FirstOrDefault();
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
            IPAddress ipAddress = GetIPsByName(hostAndPort);
            if (ipAddress != null)
            {
                string[] split = hostAndPort.Split(new string[] { ":" }, StringSplitOptions.RemoveEmptyEntries);
                hostAndPort = ipAddress.ToString() + ((split.Length == 2) ? split[1] : "");
            }
            if (!hostAndPort.Contains(":")) hostAndPort += ":8090";

            //Channel channel = new Channel(hostAndPort, channelCredentials,
            //    new ChannelOption[] { new ChannelOption("grpc.ssl_target_name_override", robotName) }
            //);
            Channel channel = new Channel(hostAndPort, ChannelCredentials.Insecure);
            try
            {
                await channel.ConnectAsync(GrpcDeadline(timeout)).ConfigureAwait(false);
            }
            catch (Exception ex)
            {
                channel?.ShutdownAsync();
                throw new EscapePodConnectionFailed("Unable to connect to Extension.", ex);
            }
            return channel;
        }

        public async Task Connect(string hostName = null)
        {
            if (this.IsConnected)
                await this.Disconnect().ConfigureAwait(false);

            this.isDisconnecting = false;
            if (!string.IsNullOrEmpty(hostName)) this.HostName = hostName;
            if (string.IsNullOrEmpty(this.HostName)) throw new ArgumentNullException("EscapePod Host must be specified", hostName);

            this.channel = await this.ConnectChannel(this.HostName, DefaultConnectionTimeout);
            this.client = new CyberVectorProxyServiceClient(channel);
            this.cancellationTokenSource = new CancellationTokenSource();

            this.eventFeed = new AsyncEventLoop<ProxyMessaage>(
                (token) => client.Subscribe(new SubscribeRequest() { KeepAlive = 20 },
                    cancellationToken: cancellationTokenSource.Token),
                this.ProcessMessage,
                this.ProcessInterrupted,
                this.ProcessException
            );
            await eventFeed.Start().ConfigureAwait(false);

            if (this.Connected != null)
                this.Connected(this, new EventArgs());
        }

        public async Task Disconnect()
        {
            try
            {
                this.isDisconnecting = true;
                try { await this.Unsubscribe().ConfigureAwait(false); }
                catch { /* suppress */ }
                this.channel?.ShutdownAsync();
            }
            catch { /* supress */ }
            finally
            {
                client = null;
                channel = null;
                eventFeed = null;

                if (this.Disconnected != null)
                    this.Disconnected(this, new EventArgs());
            }
        }

        public async Task<ProxyMessaage> Unsubscribe()
        {
            ProxyMessaage response = null;
            try
            {
                response = await this.RunMethod(client => client.UnSubscribeAsync(new UnsubscribeRequest()
                {
                    Uuid = this.uuid
                })).ConfigureAwait(false);
                this.ProcessMessage(response);
            }
            finally
            {
                this.cancellationTokenSource?.Cancel();
                this.cancellationTokenSource = null;
                this.eventFeed = null;
            }
            return response;
        }

        private void ProcessMessage(ProxyMessaage proxyMessage)
        {
            Console.WriteLine(proxyMessage);

            this.OnProxyMessageReceived?.Invoke(this, new ProxyMessageEventArgs(proxyMessage));

            switch (proxyMessage.MessageType)
            {
                case MessageType.Subscribed:
                    SubscribedMessage subscribed = new SubscribedMessage(proxyMessage.MessageData);
                    this.Subscribed?.Invoke(this, subscribed);
                    this.uuid = subscribed.Uuid;
                    break;
                case MessageType.Unsubscribed:
                    this.Unsubscribed?.Invoke(this, new EventArgs()); break;
                case MessageType.KeepAlive:
                    this.KeepAlive?.Invoke(this, new KeepAliveMessage(proxyMessage.MessageData)); break;
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
            if (!this.isDisconnecting)
            {
                _ = this.Disconnect().ConfigureAwait(false);
                throw new EscapePodConnectionFailed("EscapePod Connection Lost!");
            }
        }

        public async Task<StatusMessage> GetStatus()
        {
            StatusResponse response = await this.RunMethod(client => client.GetStatusAsync(new StatusRequest())).ConfigureAwait(false);
            return new StatusMessage(response);
        }

        public async Task<List<EscapePodIntent>> SelectIntents(string filter_json = "{}")
        {
            SelectIntentResponse response = await this.RunMethod(client => client.SelectIntentsAsync(new SelectIntentRequest()
            {
                FilterJson = filter_json
            })).ConfigureAwait(false);
            if (response.Response.Code == ResponseMessage.Types.ResponseCode.Failure)
                throw new EscapePodCommandFailed(response.Response.Message);
            return EscapePodIntent.FromResponse(response);
        }

        public async Task<string> InsertIntent(EscapePodIntent intent)
        {
            if (intent == null) throw new ArgumentNullException(nameof(intent), "Intent cannot be null!");
            InsertIntentResponse response = await this.InsertIntent(intent.ToJson()).ConfigureAwait(false);
            intent.Id = response.InsertedOid;
            return intent.Id;
        }

        public async Task<InsertIntentResponse> InsertIntent(string intentJson)
        {
            InsertIntentResponse response = await this.RunMethod(client => client.InsertIntentAsync(new InsertIntentRequest()
            {
                IntentData = intentJson
            })).ConfigureAwait(false);
            if (response.Response.Code == ResponseMessage.Types.ResponseCode.Failure)
                throw new EscapePodCommandFailed(response.Response.Message);
            return response;
        }

        public async Task<DeleteIntentResponse> DeleteIntent(string intentId)
        {
            DeleteIntentResponse response = await this.RunMethod(client => client.DeleteIntentAsync(new DeleteIntentRequest()
            {
                IntentId = intentId
            })).ConfigureAwait(false);
            if (response.Response.Code == ResponseMessage.Types.ResponseCode.Failure)
                throw new EscapePodCommandFailed(response.Response.Message);
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
            _ = this.Disconnect().ConfigureAwait(false);
        }
    }
}