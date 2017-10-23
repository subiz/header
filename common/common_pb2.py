# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: common/common.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from bitbucket.org.subiz.header.auth import auth_pb2 as bitbucket_dot_org_dot_subiz_dot_header_dot_auth_dot_auth__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='common/common.proto',
  package='common',
  syntax='proto3',
  serialized_pb=_b('\n\x13\x63ommon/common.proto\x12\x06\x63ommon\x1a*bitbucket.org/subiz/header/auth/auth.proto\"\x07\n\x05\x45mpty\"B\n\x02Id\x12\x1c\n\x03\x63tx\x18\x01 \x01(\x0b\x32\x0f.common.Context\x12\x12\n\naccount_id\x18\x02 \x01(\t\x12\n\n\x02id\x18\x03 \x01(\t\"\xd1\x01\n\x07\x43ontext\x12\x10\n\x08\x65vent_id\x18\x01 \x01(\t\x12\r\n\x05state\x18\x02 \x01(\x0c\x12\x0c\n\x04node\x18\x03 \x01(\t\x12\x13\n\x0breply_topic\x18\x04 \x01(\t\x12$\n\ncredential\x18\x06 \x01(\x0b\x32\x10.auth.Credential\x12\x0f\n\x07tracing\x18\x07 \x01(\x0c\x12\x11\n\treply_key\x18\x08 \x01(\t\x12!\n\tby_device\x18\n \x01(\x0b\x32\x0e.common.DeviceJ\x04\x08\x05\x10\x06R\x0freply_partition\"g\n\x06\x44\x65vice\x12\n\n\x02ip\x18\x03 \x01(\t\x12\x12\n\nuser_agent\x18\x04 \x01(\t\x12\x19\n\x11screen_resolution\x18\x05 \x01(\t\x12\x10\n\x08timezone\x18\x06 \x01(\t\x12\x10\n\x08language\x18\x07 \x01(\tb\x06proto3')
  ,
  dependencies=[bitbucket_dot_org_dot_subiz_dot_header_dot_auth_dot_auth__pb2.DESCRIPTOR,])




_EMPTY = _descriptor.Descriptor(
  name='Empty',
  full_name='common.Empty',
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
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=75,
  serialized_end=82,
)


_ID = _descriptor.Descriptor(
  name='Id',
  full_name='common.Id',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ctx', full_name='common.Id.ctx', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='account_id', full_name='common.Id.account_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='id', full_name='common.Id.id', index=2,
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
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=84,
  serialized_end=150,
)


_CONTEXT = _descriptor.Descriptor(
  name='Context',
  full_name='common.Context',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='event_id', full_name='common.Context.event_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='state', full_name='common.Context.state', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='node', full_name='common.Context.node', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='reply_topic', full_name='common.Context.reply_topic', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='credential', full_name='common.Context.credential', index=4,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='tracing', full_name='common.Context.tracing', index=5,
      number=7, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='reply_key', full_name='common.Context.reply_key', index=6,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='by_device', full_name='common.Context.by_device', index=7,
      number=10, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=153,
  serialized_end=362,
)


_DEVICE = _descriptor.Descriptor(
  name='Device',
  full_name='common.Device',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ip', full_name='common.Device.ip', index=0,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='user_agent', full_name='common.Device.user_agent', index=1,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='screen_resolution', full_name='common.Device.screen_resolution', index=2,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='timezone', full_name='common.Device.timezone', index=3,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='language', full_name='common.Device.language', index=4,
      number=7, type=9, cpp_type=9, label=1,
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
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=364,
  serialized_end=467,
)

_ID.fields_by_name['ctx'].message_type = _CONTEXT
_CONTEXT.fields_by_name['credential'].message_type = bitbucket_dot_org_dot_subiz_dot_header_dot_auth_dot_auth__pb2._CREDENTIAL
_CONTEXT.fields_by_name['by_device'].message_type = _DEVICE
DESCRIPTOR.message_types_by_name['Empty'] = _EMPTY
DESCRIPTOR.message_types_by_name['Id'] = _ID
DESCRIPTOR.message_types_by_name['Context'] = _CONTEXT
DESCRIPTOR.message_types_by_name['Device'] = _DEVICE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Empty = _reflection.GeneratedProtocolMessageType('Empty', (_message.Message,), dict(
  DESCRIPTOR = _EMPTY,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.Empty)
  ))
_sym_db.RegisterMessage(Empty)

Id = _reflection.GeneratedProtocolMessageType('Id', (_message.Message,), dict(
  DESCRIPTOR = _ID,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.Id)
  ))
_sym_db.RegisterMessage(Id)

Context = _reflection.GeneratedProtocolMessageType('Context', (_message.Message,), dict(
  DESCRIPTOR = _CONTEXT,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.Context)
  ))
_sym_db.RegisterMessage(Context)

Device = _reflection.GeneratedProtocolMessageType('Device', (_message.Message,), dict(
  DESCRIPTOR = _DEVICE,
  __module__ = 'common.common_pb2'
  # @@protoc_insertion_point(class_scope:common.Device)
  ))
_sym_db.RegisterMessage(Device)


# @@protoc_insertion_point(module_scope)
