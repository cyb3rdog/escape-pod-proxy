# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from . import cybervector_proxy_pb2

class CyberVectorProxyServiceStub(object):
  """The grpc-defined connection between the SDK and Cyb3rVector EscapePod Extension Proxy.
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetStatus = channel.unary_unary(
        '/cybervector.CyberVectorProxyService/GetStatus',
        request_serializer=cybervector_proxy_pb2.StatusRequest.SerializeToString,
        response_deserializer=cybervector_proxy_pb2.StatusResponse.FromString,
        )
    self.Subscribe = channel.unary_stream(
        '/cybervector.CyberVectorProxyService/Subscribe',
        request_serializer=cybervector_proxy_pb2.SubscribeRequest.SerializeToString,
        response_deserializer=cybervector_proxy_pb2.ProxyMessaage.FromString,
        )
    self.UnSubscribe = channel.unary_unary(
        '/cybervector.CyberVectorProxyService/UnSubscribe',
        request_serializer=cybervector_proxy_pb2.UnsubscribeRequest.SerializeToString,
        response_deserializer=cybervector_proxy_pb2.ProxyMessaage.FromString,
        )
    self.InsertIntent = channel.unary_unary(
        '/cybervector.CyberVectorProxyService/InsertIntent',
        request_serializer=cybervector_proxy_pb2.InsertIntentRequest.SerializeToString,
        response_deserializer=cybervector_proxy_pb2.InsertIntentResponse.FromString,
        )
    self.SelectIntents = channel.unary_unary(
        '/cybervector.CyberVectorProxyService/SelectIntents',
        request_serializer=cybervector_proxy_pb2.SelectIntentRequest.SerializeToString,
        response_deserializer=cybervector_proxy_pb2.SelectIntentResponse.FromString,
        )
    self.DeleteIntent = channel.unary_unary(
        '/cybervector.CyberVectorProxyService/DeleteIntent',
        request_serializer=cybervector_proxy_pb2.DeleteIntentRequest.SerializeToString,
        response_deserializer=cybervector_proxy_pb2.DeleteIntentResponse.FromString,
        )


class CyberVectorProxyService(object):
  """The grpc-defined connection between the SDK and Cyb3rVector EscapePod Extension Proxy.
  """
  def GetStatus(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Subscribe(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def UnSubscribe(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def InsertIntent(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SelectIntents(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteIntent(self, request, context):
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_CyberVectorProxyService_to_server(servicer, server):
  rpc_method_handlers = {
      'GetStatus': grpc.unary_unary_rpc_method_handler(
          servicer.GetStatus,
          request_deserializer=cybervector_proxy_pb2.StatusRequest.FromString,
          response_serializer=cybervector_proxy_pb2.StatusResponse.SerializeToString,
      ),
      'Subscribe': grpc.unary_stream_rpc_method_handler(
          servicer.Subscribe,
          request_deserializer=cybervector_proxy_pb2.SubscribeRequest.FromString,
          response_serializer=cybervector_proxy_pb2.ProxyMessaage.SerializeToString,
      ),
      'UnSubscribe': grpc.unary_unary_rpc_method_handler(
          servicer.UnSubscribe,
          request_deserializer=cybervector_proxy_pb2.UnsubscribeRequest.FromString,
          response_serializer=cybervector_proxy_pb2.ProxyMessaage.SerializeToString,
      ),
      'InsertIntent': grpc.unary_unary_rpc_method_handler(
          servicer.InsertIntent,
          request_deserializer=cybervector_proxy_pb2.InsertIntentRequest.FromString,
          response_serializer=cybervector_proxy_pb2.InsertIntentResponse.SerializeToString,
      ),
      'SelectIntents': grpc.unary_unary_rpc_method_handler(
          servicer.SelectIntents,
          request_deserializer=cybervector_proxy_pb2.SelectIntentRequest.FromString,
          response_serializer=cybervector_proxy_pb2.SelectIntentResponse.SerializeToString,
      ),
      'DeleteIntent': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteIntent,
          request_deserializer=cybervector_proxy_pb2.DeleteIntentRequest.FromString,
          response_serializer=cybervector_proxy_pb2.DeleteIntentResponse.SerializeToString,
      )
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'cybervector.CyberVectorProxyService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
