# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/api.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='api/api.proto',
  package='api',
  syntax='proto2',
  serialized_pb=_b('\n\rapi/api.proto\x12\x03\x61pi\"\x07\n\x05\x45mpty\"!\n\x10SendMessageState\x12\r\n\x05topic\x18\x01 \x01(\t\"\x16\n\x05State\x12\r\n\x05\x65vent\x18\x03 \x01(\t*;\n\x06Method\x12\x07\n\x03GET\x10\x00\x12\x08\n\x04POST\x10\x01\x12\x07\n\x03PUT\x10\x02\x12\n\n\x06\x44\x45LETE\x10\x03\x12\t\n\x05PATCH\x10\x04*,\n\x05\x45vent\x12\x15\n\x11\x43onversationStart\x10\x02\x12\x0c\n\x08\x41piReply\x10\x04')
)

_METHOD = _descriptor.EnumDescriptor(
  name='Method',
  full_name='api.Method',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='GET', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='POST', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PUT', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DELETE', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PATCH', index=4, number=4,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=90,
  serialized_end=149,
)
_sym_db.RegisterEnumDescriptor(_METHOD)

Method = enum_type_wrapper.EnumTypeWrapper(_METHOD)
_EVENT = _descriptor.EnumDescriptor(
  name='Event',
  full_name='api.Event',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ConversationStart', index=0, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ApiReply', index=1, number=4,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=151,
  serialized_end=195,
)
_sym_db.RegisterEnumDescriptor(_EVENT)

Event = enum_type_wrapper.EnumTypeWrapper(_EVENT)
GET = 0
POST = 1
PUT = 2
DELETE = 3
PATCH = 4
ConversationStart = 2
ApiReply = 4



_EMPTY = _descriptor.Descriptor(
  name='Empty',
  full_name='api.Empty',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=22,
  serialized_end=29,
)


_SENDMESSAGESTATE = _descriptor.Descriptor(
  name='SendMessageState',
  full_name='api.SendMessageState',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='topic', full_name='api.SendMessageState.topic', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=31,
  serialized_end=64,
)


_STATE = _descriptor.Descriptor(
  name='State',
  full_name='api.State',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='event', full_name='api.State.event', index=0,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=66,
  serialized_end=88,
)

DESCRIPTOR.message_types_by_name['Empty'] = _EMPTY
DESCRIPTOR.message_types_by_name['SendMessageState'] = _SENDMESSAGESTATE
DESCRIPTOR.message_types_by_name['State'] = _STATE
DESCRIPTOR.enum_types_by_name['Method'] = _METHOD
DESCRIPTOR.enum_types_by_name['Event'] = _EVENT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Empty = _reflection.GeneratedProtocolMessageType('Empty', (_message.Message,), dict(
  DESCRIPTOR = _EMPTY,
  __module__ = 'api.api_pb2'
  # @@protoc_insertion_point(class_scope:api.Empty)
  ))
_sym_db.RegisterMessage(Empty)

SendMessageState = _reflection.GeneratedProtocolMessageType('SendMessageState', (_message.Message,), dict(
  DESCRIPTOR = _SENDMESSAGESTATE,
  __module__ = 'api.api_pb2'
  # @@protoc_insertion_point(class_scope:api.SendMessageState)
  ))
_sym_db.RegisterMessage(SendMessageState)

State = _reflection.GeneratedProtocolMessageType('State', (_message.Message,), dict(
  DESCRIPTOR = _STATE,
  __module__ = 'api.api_pb2'
  # @@protoc_insertion_point(class_scope:api.State)
  ))
_sym_db.RegisterMessage(State)


# @@protoc_insertion_point(module_scope)
