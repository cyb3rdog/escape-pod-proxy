# Generated by the Protocol Buffers compiler. DO NOT EDIT!
# source: cybervector_proxy.proto
# plugin: grpclib.plugin.main
import abc
import typing

import grpclib.const
import grpclib.client
if typing.TYPE_CHECKING:
    import grpclib.server

import cybervector_proxy_pb2


class CyberVectorProxyServiceBase(abc.ABC):

    @abc.abstractmethod
    async def GetStatus(self, stream: 'grpclib.server.Stream[cybervector_proxy_pb2.StatusRequest, cybervector_proxy_pb2.StatusResponse]') -> None:
        pass

    @abc.abstractmethod
    async def Subscribe(self, stream: 'grpclib.server.Stream[cybervector_proxy_pb2.SubscribeRequest, cybervector_proxy_pb2.ProxyMessaage]') -> None:
        pass

    @abc.abstractmethod
    async def UnSubscribe(self, stream: 'grpclib.server.Stream[cybervector_proxy_pb2.UnsubscribeRequest, cybervector_proxy_pb2.ProxyMessaage]') -> None:
        pass

    @abc.abstractmethod
    async def InsertIntent(self, stream: 'grpclib.server.Stream[cybervector_proxy_pb2.InsertIntentRequest, cybervector_proxy_pb2.InsertIntentResponse]') -> None:
        pass

    @abc.abstractmethod
    async def SelectIntents(self, stream: 'grpclib.server.Stream[cybervector_proxy_pb2.SelectIntentRequest, cybervector_proxy_pb2.SelectIntentResponse]') -> None:
        pass

    @abc.abstractmethod
    async def DeleteIntent(self, stream: 'grpclib.server.Stream[cybervector_proxy_pb2.DeleteIntentRequest, cybervector_proxy_pb2.DeleteIntentResponse]') -> None:
        pass

    def __mapping__(self) -> typing.Dict[str, grpclib.const.Handler]:
        return {
            '/cybervector.CyberVectorProxyService/GetStatus': grpclib.const.Handler(
                self.GetStatus,
                grpclib.const.Cardinality.UNARY_UNARY,
                cybervector_proxy_pb2.StatusRequest,
                cybervector_proxy_pb2.StatusResponse,
            ),
            '/cybervector.CyberVectorProxyService/Subscribe': grpclib.const.Handler(
                self.Subscribe,
                grpclib.const.Cardinality.UNARY_STREAM,
                cybervector_proxy_pb2.SubscribeRequest,
                cybervector_proxy_pb2.ProxyMessaage,
            ),
            '/cybervector.CyberVectorProxyService/UnSubscribe': grpclib.const.Handler(
                self.UnSubscribe,
                grpclib.const.Cardinality.UNARY_UNARY,
                cybervector_proxy_pb2.UnsubscribeRequest,
                cybervector_proxy_pb2.ProxyMessaage,
            ),
            '/cybervector.CyberVectorProxyService/InsertIntent': grpclib.const.Handler(
                self.InsertIntent,
                grpclib.const.Cardinality.UNARY_UNARY,
                cybervector_proxy_pb2.InsertIntentRequest,
                cybervector_proxy_pb2.InsertIntentResponse,
            ),
            '/cybervector.CyberVectorProxyService/SelectIntents': grpclib.const.Handler(
                self.SelectIntents,
                grpclib.const.Cardinality.UNARY_UNARY,
                cybervector_proxy_pb2.SelectIntentRequest,
                cybervector_proxy_pb2.SelectIntentResponse,
            ),
            '/cybervector.CyberVectorProxyService/DeleteIntent': grpclib.const.Handler(
                self.DeleteIntent,
                grpclib.const.Cardinality.UNARY_UNARY,
                cybervector_proxy_pb2.DeleteIntentRequest,
                cybervector_proxy_pb2.DeleteIntentResponse,
            ),
        }


class CyberVectorProxyServiceStub:

    def __init__(self, channel: grpclib.client.Channel) -> None:
        self.GetStatus = grpclib.client.UnaryUnaryMethod(
            channel,
            '/cybervector.CyberVectorProxyService/GetStatus',
            cybervector_proxy_pb2.StatusRequest,
            cybervector_proxy_pb2.StatusResponse,
        )
        self.Subscribe = grpclib.client.UnaryStreamMethod(
            channel,
            '/cybervector.CyberVectorProxyService/Subscribe',
            cybervector_proxy_pb2.SubscribeRequest,
            cybervector_proxy_pb2.ProxyMessaage,
        )
        self.UnSubscribe = grpclib.client.UnaryUnaryMethod(
            channel,
            '/cybervector.CyberVectorProxyService/UnSubscribe',
            cybervector_proxy_pb2.UnsubscribeRequest,
            cybervector_proxy_pb2.ProxyMessaage,
        )
        self.InsertIntent = grpclib.client.UnaryUnaryMethod(
            channel,
            '/cybervector.CyberVectorProxyService/InsertIntent',
            cybervector_proxy_pb2.InsertIntentRequest,
            cybervector_proxy_pb2.InsertIntentResponse,
        )
        self.SelectIntents = grpclib.client.UnaryUnaryMethod(
            channel,
            '/cybervector.CyberVectorProxyService/SelectIntents',
            cybervector_proxy_pb2.SelectIntentRequest,
            cybervector_proxy_pb2.SelectIntentResponse,
        )
        self.DeleteIntent = grpclib.client.UnaryUnaryMethod(
            channel,
            '/cybervector.CyberVectorProxyService/DeleteIntent',
            cybervector_proxy_pb2.DeleteIntentRequest,
            cybervector_proxy_pb2.DeleteIntentResponse,
        )