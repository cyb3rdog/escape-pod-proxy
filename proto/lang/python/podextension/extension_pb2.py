# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: extension.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='extension.proto',
  package='podextension',
  syntax='proto3',
  serialized_options=_b('Z?github.com/cyb3rdog/escape-pod-proxy/proto/lang/go/podextension'),
  serialized_pb=_b('\n\x0f\x65xtension.proto\x12\x0cpodextension\"\x8e\x01\n\x08UnaryReq\x12\x13\n\x0bintent_name\x18\x01 \x01(\t\x12:\n\nparameters\x18\x02 \x03(\x0b\x32&.podextension.UnaryReq.ParametersEntry\x1a\x31\n\x0fParametersEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\x90\x01\n\tUnaryResp\x12\x13\n\x0bintent_name\x18\x01 \x01(\t\x12;\n\nparameters\x18\x02 \x03(\x0b\x32\'.podextension.UnaryResp.ParametersEntry\x1a\x31\n\x0fParametersEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x32H\n\x0cPodExtension\x12\x38\n\x05Unary\x12\x16.podextension.UnaryReq\x1a\x17.podextension.UnaryRespBAZ?github.com/cyb3rdog/escape-pod-proxy/proto/lang/go/podextensionb\x06proto3')
)




_UNARYREQ_PARAMETERSENTRY = _descriptor.Descriptor(
  name='ParametersEntry',
  full_name='podextension.UnaryReq.ParametersEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='podextension.UnaryReq.ParametersEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='podextension.UnaryReq.ParametersEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=127,
  serialized_end=176,
)

_UNARYREQ = _descriptor.Descriptor(
  name='UnaryReq',
  full_name='podextension.UnaryReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='intent_name', full_name='podextension.UnaryReq.intent_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parameters', full_name='podextension.UnaryReq.parameters', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_UNARYREQ_PARAMETERSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=34,
  serialized_end=176,
)


_UNARYRESP_PARAMETERSENTRY = _descriptor.Descriptor(
  name='ParametersEntry',
  full_name='podextension.UnaryResp.ParametersEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='podextension.UnaryResp.ParametersEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='podextension.UnaryResp.ParametersEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=127,
  serialized_end=176,
)

_UNARYRESP = _descriptor.Descriptor(
  name='UnaryResp',
  full_name='podextension.UnaryResp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='intent_name', full_name='podextension.UnaryResp.intent_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parameters', full_name='podextension.UnaryResp.parameters', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_UNARYRESP_PARAMETERSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=179,
  serialized_end=323,
)

_UNARYREQ_PARAMETERSENTRY.containing_type = _UNARYREQ
_UNARYREQ.fields_by_name['parameters'].message_type = _UNARYREQ_PARAMETERSENTRY
_UNARYRESP_PARAMETERSENTRY.containing_type = _UNARYRESP
_UNARYRESP.fields_by_name['parameters'].message_type = _UNARYRESP_PARAMETERSENTRY
DESCRIPTOR.message_types_by_name['UnaryReq'] = _UNARYREQ
DESCRIPTOR.message_types_by_name['UnaryResp'] = _UNARYRESP
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

UnaryReq = _reflection.GeneratedProtocolMessageType('UnaryReq', (_message.Message,), dict(

  ParametersEntry = _reflection.GeneratedProtocolMessageType('ParametersEntry', (_message.Message,), dict(
    DESCRIPTOR = _UNARYREQ_PARAMETERSENTRY,
    __module__ = 'extension_pb2'
    # @@protoc_insertion_point(class_scope:podextension.UnaryReq.ParametersEntry)
    ))
  ,
  DESCRIPTOR = _UNARYREQ,
  __module__ = 'extension_pb2'
  # @@protoc_insertion_point(class_scope:podextension.UnaryReq)
  ))
_sym_db.RegisterMessage(UnaryReq)
_sym_db.RegisterMessage(UnaryReq.ParametersEntry)

UnaryResp = _reflection.GeneratedProtocolMessageType('UnaryResp', (_message.Message,), dict(

  ParametersEntry = _reflection.GeneratedProtocolMessageType('ParametersEntry', (_message.Message,), dict(
    DESCRIPTOR = _UNARYRESP_PARAMETERSENTRY,
    __module__ = 'extension_pb2'
    # @@protoc_insertion_point(class_scope:podextension.UnaryResp.ParametersEntry)
    ))
  ,
  DESCRIPTOR = _UNARYRESP,
  __module__ = 'extension_pb2'
  # @@protoc_insertion_point(class_scope:podextension.UnaryResp)
  ))
_sym_db.RegisterMessage(UnaryResp)
_sym_db.RegisterMessage(UnaryResp.ParametersEntry)


DESCRIPTOR._options = None
_UNARYREQ_PARAMETERSENTRY._options = None
_UNARYRESP_PARAMETERSENTRY._options = None

_PODEXTENSION = _descriptor.ServiceDescriptor(
  name='PodExtension',
  full_name='podextension.PodExtension',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=325,
  serialized_end=397,
  methods=[
  _descriptor.MethodDescriptor(
    name='Unary',
    full_name='podextension.PodExtension.Unary',
    index=0,
    containing_service=None,
    input_type=_UNARYREQ,
    output_type=_UNARYRESP,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_PODEXTENSION)

DESCRIPTOR.services_by_name['PodExtension'] = _PODEXTENSION

# @@protoc_insertion_point(module_scope)
