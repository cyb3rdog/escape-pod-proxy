# Generated by the Protocol Buffers compiler. DO NOT EDIT!
# source: extension.proto
# plugin: grpclib.plugin.main
import abc
import typing

import grpclib.const
import grpclib.client
if typing.TYPE_CHECKING:
    import grpclib.server

import extension_pb2


class PodExtensionBase(abc.ABC):

    @abc.abstractmethod
    async def Unary(self, stream: 'grpclib.server.Stream[extension_pb2.UnaryReq, extension_pb2.UnaryResp]') -> None:
        pass

    def __mapping__(self) -> typing.Dict[str, grpclib.const.Handler]:
        return {
            '/podextension.PodExtension/Unary': grpclib.const.Handler(
                self.Unary,
                grpclib.const.Cardinality.UNARY_UNARY,
                extension_pb2.UnaryReq,
                extension_pb2.UnaryResp,
            ),
        }


class PodExtensionStub:

    def __init__(self, channel: grpclib.client.Channel) -> None:
        self.Unary = grpclib.client.UnaryUnaryMethod(
            channel,
            '/podextension.PodExtension/Unary',
            extension_pb2.UnaryReq,
            extension_pb2.UnaryResp,
        )
