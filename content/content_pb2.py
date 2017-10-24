# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: content/content.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from bitbucket.org.subiz.header.common import common_pb2 as bitbucket_dot_org_dot_subiz_dot_header_dot_common_dot_common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='content/content.proto',
  package='content',
  syntax='proto2',
  serialized_pb=_b('\n\x15\x63ontent/content.proto\x12\x07\x63ontent\x1a.bitbucket.org/subiz/header/common/common.proto\"&\n\x08KeyValue\x12\x0b\n\x03key\x18\x03 \x01(\t\x12\r\n\x05value\x18\x04 \x01(\t\"\xd5\x02\n\x07\x43ontent\x12\x1c\n\x03\x63tx\x18\x01 \x01(\x0b\x32\x0f.common.Context\x12\n\n\x02id\x18\x02 \x01(\t\x12\x12\n\naccount_id\x18\x03 \x01(\t\x12\x0c\n\x04type\x18\x04 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x05 \x01(\t\x12\r\n\x05title\x18\x06 \x01(\t\x12\x0c\n\x04link\x18\x07 \x01(\t\x12\x0e\n\x06labels\x18\x08 \x03(\t\x12\x14\n\x0c\x61vailability\x18\t \x01(\x08\x12\r\n\x05price\x18\n \x01(\t\x12\x10\n\x08\x63urrency\x18\x0b \x01(\t\x12\x12\n\nsale_price\x18\x0c \x01(\t\x12!\n\x06\x66ields\x18\r \x03(\x0b\x32\x11.content.KeyValue\x12\x12\n\ncategories\x18\x0e \x03(\t\x12\"\n\x08relateds\x18\x0f \x03(\x0b\x32\x10.content.Content\x12\x16\n\x0e\x61ttachment_url\x18\x10 \x01(\t\".\n\x08\x43ontents\x12\"\n\x08\x63ontents\x18\x01 \x03(\x0b\x32\x10.content.Content\"\x07\n\x05\x45mpty\"T\n\x0bListRequest\x12\x12\n\naccount_id\x18\x06 \x01(\t\x12\x10\n\x08start_id\x18\x02 \x01(\t\x12\x10\n\x08\x63\x61tegory\x18\x03 \x01(\t\x12\r\n\x05limit\x18\x04 \x01(\x05')
  ,
  dependencies=[bitbucket_dot_org_dot_subiz_dot_header_dot_common_dot_common__pb2.DESCRIPTOR,])




_KEYVALUE = _descriptor.Descriptor(
  name='KeyValue',
  full_name='content.KeyValue',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='content.KeyValue.key', index=0,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='value', full_name='content.KeyValue.value', index=1,
      number=4, type=9, cpp_type=9, label=1,
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
  serialized_start=82,
  serialized_end=120,
)


_CONTENT = _descriptor.Descriptor(
  name='Content',
  full_name='content.Content',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ctx', full_name='content.Content.ctx', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='id', full_name='content.Content.id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='account_id', full_name='content.Content.account_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='type', full_name='content.Content.type', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='description', full_name='content.Content.description', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='title', full_name='content.Content.title', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='link', full_name='content.Content.link', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='labels', full_name='content.Content.labels', index=7,
      number=8, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='availability', full_name='content.Content.availability', index=8,
      number=9, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='price', full_name='content.Content.price', index=9,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='currency', full_name='content.Content.currency', index=10,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='sale_price', full_name='content.Content.sale_price', index=11,
      number=12, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='fields', full_name='content.Content.fields', index=12,
      number=13, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='categories', full_name='content.Content.categories', index=13,
      number=14, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='relateds', full_name='content.Content.relateds', index=14,
      number=15, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='attachment_url', full_name='content.Content.attachment_url', index=15,
      number=16, type=9, cpp_type=9, label=1,
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
  serialized_start=123,
  serialized_end=464,
)


_CONTENTS = _descriptor.Descriptor(
  name='Contents',
  full_name='content.Contents',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='contents', full_name='content.Contents.contents', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=466,
  serialized_end=512,
)


_EMPTY = _descriptor.Descriptor(
  name='Empty',
  full_name='content.Empty',
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
  serialized_start=514,
  serialized_end=521,
)


_LISTREQUEST = _descriptor.Descriptor(
  name='ListRequest',
  full_name='content.ListRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='account_id', full_name='content.ListRequest.account_id', index=0,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='start_id', full_name='content.ListRequest.start_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='category', full_name='content.ListRequest.category', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='limit', full_name='content.ListRequest.limit', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=523,
  serialized_end=607,
)

_CONTENT.fields_by_name['ctx'].message_type = bitbucket_dot_org_dot_subiz_dot_header_dot_common_dot_common__pb2._CONTEXT
_CONTENT.fields_by_name['fields'].message_type = _KEYVALUE
_CONTENT.fields_by_name['relateds'].message_type = _CONTENT
_CONTENTS.fields_by_name['contents'].message_type = _CONTENT
DESCRIPTOR.message_types_by_name['KeyValue'] = _KEYVALUE
DESCRIPTOR.message_types_by_name['Content'] = _CONTENT
DESCRIPTOR.message_types_by_name['Contents'] = _CONTENTS
DESCRIPTOR.message_types_by_name['Empty'] = _EMPTY
DESCRIPTOR.message_types_by_name['ListRequest'] = _LISTREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

KeyValue = _reflection.GeneratedProtocolMessageType('KeyValue', (_message.Message,), dict(
  DESCRIPTOR = _KEYVALUE,
  __module__ = 'content.content_pb2'
  # @@protoc_insertion_point(class_scope:content.KeyValue)
  ))
_sym_db.RegisterMessage(KeyValue)

Content = _reflection.GeneratedProtocolMessageType('Content', (_message.Message,), dict(
  DESCRIPTOR = _CONTENT,
  __module__ = 'content.content_pb2'
  # @@protoc_insertion_point(class_scope:content.Content)
  ))
_sym_db.RegisterMessage(Content)

Contents = _reflection.GeneratedProtocolMessageType('Contents', (_message.Message,), dict(
  DESCRIPTOR = _CONTENTS,
  __module__ = 'content.content_pb2'
  # @@protoc_insertion_point(class_scope:content.Contents)
  ))
_sym_db.RegisterMessage(Contents)

Empty = _reflection.GeneratedProtocolMessageType('Empty', (_message.Message,), dict(
  DESCRIPTOR = _EMPTY,
  __module__ = 'content.content_pb2'
  # @@protoc_insertion_point(class_scope:content.Empty)
  ))
_sym_db.RegisterMessage(Empty)

ListRequest = _reflection.GeneratedProtocolMessageType('ListRequest', (_message.Message,), dict(
  DESCRIPTOR = _LISTREQUEST,
  __module__ = 'content.content_pb2'
  # @@protoc_insertion_point(class_scope:content.ListRequest)
  ))
_sym_db.RegisterMessage(ListRequest)


# @@protoc_insertion_point(module_scope)
