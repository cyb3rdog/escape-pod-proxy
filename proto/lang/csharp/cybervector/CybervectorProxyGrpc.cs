// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: cybervector_proxy.proto
// </auto-generated>
#pragma warning disable 0414, 1591
#region Designer generated code

using grpc = global::Grpc.Core;

namespace Cybervector {
  public static partial class CyberVectorProxyService
  {
    static readonly string __ServiceName = "cybervector.CyberVectorProxyService";

    static readonly grpc::Marshaller<global::Cybervector.StatusRequest> __Marshaller_cybervector_StatusRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Cybervector.StatusRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Cybervector.StatusResponse> __Marshaller_cybervector_StatusResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Cybervector.StatusResponse.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Cybervector.SubscribeRequest> __Marshaller_cybervector_SubscribeRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Cybervector.SubscribeRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Cybervector.ProxyMessaage> __Marshaller_cybervector_ProxyMessaage = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Cybervector.ProxyMessaage.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Cybervector.UnsubscribeRequest> __Marshaller_cybervector_UnsubscribeRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Cybervector.UnsubscribeRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Cybervector.InsertIntentRequest> __Marshaller_cybervector_InsertIntentRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Cybervector.InsertIntentRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Cybervector.InsertIntentResponse> __Marshaller_cybervector_InsertIntentResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Cybervector.InsertIntentResponse.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Cybervector.SelectIntentRequest> __Marshaller_cybervector_SelectIntentRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Cybervector.SelectIntentRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Cybervector.SelectIntentResponse> __Marshaller_cybervector_SelectIntentResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Cybervector.SelectIntentResponse.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Cybervector.DeleteIntentRequest> __Marshaller_cybervector_DeleteIntentRequest = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Cybervector.DeleteIntentRequest.Parser.ParseFrom);
    static readonly grpc::Marshaller<global::Cybervector.DeleteIntentResponse> __Marshaller_cybervector_DeleteIntentResponse = grpc::Marshallers.Create((arg) => global::Google.Protobuf.MessageExtensions.ToByteArray(arg), global::Cybervector.DeleteIntentResponse.Parser.ParseFrom);

    static readonly grpc::Method<global::Cybervector.StatusRequest, global::Cybervector.StatusResponse> __Method_GetStatus = new grpc::Method<global::Cybervector.StatusRequest, global::Cybervector.StatusResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetStatus",
        __Marshaller_cybervector_StatusRequest,
        __Marshaller_cybervector_StatusResponse);

    static readonly grpc::Method<global::Cybervector.SubscribeRequest, global::Cybervector.ProxyMessaage> __Method_Subscribe = new grpc::Method<global::Cybervector.SubscribeRequest, global::Cybervector.ProxyMessaage>(
        grpc::MethodType.ServerStreaming,
        __ServiceName,
        "Subscribe",
        __Marshaller_cybervector_SubscribeRequest,
        __Marshaller_cybervector_ProxyMessaage);

    static readonly grpc::Method<global::Cybervector.UnsubscribeRequest, global::Cybervector.ProxyMessaage> __Method_UnSubscribe = new grpc::Method<global::Cybervector.UnsubscribeRequest, global::Cybervector.ProxyMessaage>(
        grpc::MethodType.Unary,
        __ServiceName,
        "UnSubscribe",
        __Marshaller_cybervector_UnsubscribeRequest,
        __Marshaller_cybervector_ProxyMessaage);

    static readonly grpc::Method<global::Cybervector.InsertIntentRequest, global::Cybervector.InsertIntentResponse> __Method_InsertIntent = new grpc::Method<global::Cybervector.InsertIntentRequest, global::Cybervector.InsertIntentResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "InsertIntent",
        __Marshaller_cybervector_InsertIntentRequest,
        __Marshaller_cybervector_InsertIntentResponse);

    static readonly grpc::Method<global::Cybervector.SelectIntentRequest, global::Cybervector.SelectIntentResponse> __Method_SelectIntents = new grpc::Method<global::Cybervector.SelectIntentRequest, global::Cybervector.SelectIntentResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "SelectIntents",
        __Marshaller_cybervector_SelectIntentRequest,
        __Marshaller_cybervector_SelectIntentResponse);

    static readonly grpc::Method<global::Cybervector.DeleteIntentRequest, global::Cybervector.DeleteIntentResponse> __Method_DeleteIntent = new grpc::Method<global::Cybervector.DeleteIntentRequest, global::Cybervector.DeleteIntentResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "DeleteIntent",
        __Marshaller_cybervector_DeleteIntentRequest,
        __Marshaller_cybervector_DeleteIntentResponse);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Cybervector.CybervectorProxyReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of CyberVectorProxyService</summary>
    public abstract partial class CyberVectorProxyServiceBase
    {
      public virtual global::System.Threading.Tasks.Task<global::Cybervector.StatusResponse> GetStatus(global::Cybervector.StatusRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task Subscribe(global::Cybervector.SubscribeRequest request, grpc::IServerStreamWriter<global::Cybervector.ProxyMessaage> responseStream, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Cybervector.ProxyMessaage> UnSubscribe(global::Cybervector.UnsubscribeRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Cybervector.InsertIntentResponse> InsertIntent(global::Cybervector.InsertIntentRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Cybervector.SelectIntentResponse> SelectIntents(global::Cybervector.SelectIntentRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      public virtual global::System.Threading.Tasks.Task<global::Cybervector.DeleteIntentResponse> DeleteIntent(global::Cybervector.DeleteIntentRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for CyberVectorProxyService</summary>
    public partial class CyberVectorProxyServiceClient : grpc::ClientBase<CyberVectorProxyServiceClient>
    {
      /// <summary>Creates a new client for CyberVectorProxyService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      public CyberVectorProxyServiceClient(grpc::Channel channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for CyberVectorProxyService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      public CyberVectorProxyServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      protected CyberVectorProxyServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      protected CyberVectorProxyServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      public virtual global::Cybervector.StatusResponse GetStatus(global::Cybervector.StatusRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetStatus(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Cybervector.StatusResponse GetStatus(global::Cybervector.StatusRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetStatus, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Cybervector.StatusResponse> GetStatusAsync(global::Cybervector.StatusRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetStatusAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Cybervector.StatusResponse> GetStatusAsync(global::Cybervector.StatusRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetStatus, null, options, request);
      }
      public virtual grpc::AsyncServerStreamingCall<global::Cybervector.ProxyMessaage> Subscribe(global::Cybervector.SubscribeRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return Subscribe(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncServerStreamingCall<global::Cybervector.ProxyMessaage> Subscribe(global::Cybervector.SubscribeRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncServerStreamingCall(__Method_Subscribe, null, options, request);
      }
      public virtual global::Cybervector.ProxyMessaage UnSubscribe(global::Cybervector.UnsubscribeRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return UnSubscribe(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Cybervector.ProxyMessaage UnSubscribe(global::Cybervector.UnsubscribeRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_UnSubscribe, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Cybervector.ProxyMessaage> UnSubscribeAsync(global::Cybervector.UnsubscribeRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return UnSubscribeAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Cybervector.ProxyMessaage> UnSubscribeAsync(global::Cybervector.UnsubscribeRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_UnSubscribe, null, options, request);
      }
      public virtual global::Cybervector.InsertIntentResponse InsertIntent(global::Cybervector.InsertIntentRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return InsertIntent(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Cybervector.InsertIntentResponse InsertIntent(global::Cybervector.InsertIntentRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_InsertIntent, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Cybervector.InsertIntentResponse> InsertIntentAsync(global::Cybervector.InsertIntentRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return InsertIntentAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Cybervector.InsertIntentResponse> InsertIntentAsync(global::Cybervector.InsertIntentRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_InsertIntent, null, options, request);
      }
      public virtual global::Cybervector.SelectIntentResponse SelectIntents(global::Cybervector.SelectIntentRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return SelectIntents(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Cybervector.SelectIntentResponse SelectIntents(global::Cybervector.SelectIntentRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_SelectIntents, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Cybervector.SelectIntentResponse> SelectIntentsAsync(global::Cybervector.SelectIntentRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return SelectIntentsAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Cybervector.SelectIntentResponse> SelectIntentsAsync(global::Cybervector.SelectIntentRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_SelectIntents, null, options, request);
      }
      public virtual global::Cybervector.DeleteIntentResponse DeleteIntent(global::Cybervector.DeleteIntentRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return DeleteIntent(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual global::Cybervector.DeleteIntentResponse DeleteIntent(global::Cybervector.DeleteIntentRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_DeleteIntent, null, options, request);
      }
      public virtual grpc::AsyncUnaryCall<global::Cybervector.DeleteIntentResponse> DeleteIntentAsync(global::Cybervector.DeleteIntentRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return DeleteIntentAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      public virtual grpc::AsyncUnaryCall<global::Cybervector.DeleteIntentResponse> DeleteIntentAsync(global::Cybervector.DeleteIntentRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_DeleteIntent, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      protected override CyberVectorProxyServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new CyberVectorProxyServiceClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    public static grpc::ServerServiceDefinition BindService(CyberVectorProxyServiceBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_GetStatus, serviceImpl.GetStatus)
          .AddMethod(__Method_Subscribe, serviceImpl.Subscribe)
          .AddMethod(__Method_UnSubscribe, serviceImpl.UnSubscribe)
          .AddMethod(__Method_InsertIntent, serviceImpl.InsertIntent)
          .AddMethod(__Method_SelectIntents, serviceImpl.SelectIntents)
          .AddMethod(__Method_DeleteIntent, serviceImpl.DeleteIntent).Build();
    }

    /// <summary>Register service method implementations with a service binder. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    public static void BindService(grpc::ServiceBinderBase serviceBinder, CyberVectorProxyServiceBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_GetStatus, serviceImpl.GetStatus);
      serviceBinder.AddMethod(__Method_Subscribe, serviceImpl.Subscribe);
      serviceBinder.AddMethod(__Method_UnSubscribe, serviceImpl.UnSubscribe);
      serviceBinder.AddMethod(__Method_InsertIntent, serviceImpl.InsertIntent);
      serviceBinder.AddMethod(__Method_SelectIntents, serviceImpl.SelectIntents);
      serviceBinder.AddMethod(__Method_DeleteIntent, serviceImpl.DeleteIntent);
    }

  }
}
#endregion