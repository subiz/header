# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: event/event.proto

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


from bitbucket.org.subiz.header.common import common_pb2 as bitbucket_dot_org_dot_subiz_dot_header_dot_common_dot_common__pb2
from bitbucket.org.subiz.header.lang import lang_pb2 as bitbucket_dot_org_dot_subiz_dot_header_dot_lang_dot_lang__pb2
from bitbucket.org.subiz.header.account import account_pb2 as bitbucket_dot_org_dot_subiz_dot_header_dot_account_dot_account__pb2
from bitbucket.org.subiz.header.conversation import conversation_pb2 as bitbucket_dot_org_dot_subiz_dot_header_dot_conversation_dot_conversation__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='event/event.proto',
  package='event',
  syntax='proto2',
  serialized_pb=_b('\n\x11\x65vent/event.proto\x12\x05\x65vent\x1a.bitbucket.org/subiz/header/common/common.proto\x1a*bitbucket.org/subiz/header/lang/lang.proto\x1a\x30\x62itbucket.org/subiz/header/account/account.proto\x1a:bitbucket.org/subiz/header/conversation/conversation.proto\"f\n\x16RawEventCreatedPayload\x12\x1c\n\x03\x63tx\x18\x01 \x01(\x0b\x32\x0f.common.Context\x12\x0e\n\x06sub_id\x18\x03 \x01(\t\x12\r\n\x05topic\x18\x04 \x01(\t\x12\x0f\n\x07payload\x18\x05 \x01(\t\"3\n\x12UnsubscribeMessage\x12\r\n\x05topic\x18\x03 \x01(\t\x12\x0e\n\x06sub_id\x18\x04 \x01(\t\",\n\tRawEvents\x12\x1f\n\x06\x65vents\x18\x04 \x03(\x0b\x32\x0f.event.RawEvent\"J\n\x02\x42y\x12\x1e\n\x06\x64\x65vice\x18\x02 \x01(\x0b\x32\x0e.common.Device\x12\x11\n\tmember_id\x18\x03 \x01(\t\x12\x11\n\tclient_id\x18\x04 \x01(\t\"\xa9\x05\n\x08RawEvent\x12\x1c\n\x03\x63tx\x18\x01 \x01(\x0b\x32\x0f.common.Context\x12\n\n\x02id\x18\x03 \x01(\t\x12\x12\n\naccount_id\x18\x04 \x01(\t\x12\x0f\n\x07\x63reated\x18\x08 \x01(\x03\x12\x0c\n\x04type\x18\t \x01(\t\x12\x0e\n\x06topics\x18\x0b \x03(\t\x12\x15\n\x02\x62y\x18\x32 \x01(\x0b\x32\t.event.By\x12\x0e\n\x06object\x18\r \x01(\t\x12\x17\n\x0f\x63onversation_id\x18\x14 \x01(\t\x12#\n\x07\x61\x63\x63ount\x18& \x01(\x0b\x32\x10.account.AccountH\x00\x12\x1f\n\x05\x61gent\x18$ \x01(\x0b\x32\x0e.account.AgentH\x00\x12(\n\x07message\x18% \x01(\x0b\x32\x15.conversation.MessageH\x00\x12\x32\n\x0c\x63onversation\x18\' \x01(\x0b\x32\x1a.conversation.ConversationH\x00\x12*\n\x08postback\x18( \x01(\x0b\x32\x16.conversation.PostbackH\x00\x42\x06\n\x04\x64\x61taJ\x04\x08\x10\x10\x11J\x04\x08\x11\x10\x12J\x04\x08\x12\x10\x13J\x04\x08\x13\x10\x14J\x04\x08\x15\x10\x16J\x04\x08\x16\x10\x17J\x04\x08\x17\x10\x18J\x04\x08\x18\x10\x19J\x04\x08\x19\x10\x1aJ\x04\x08\x1a\x10\x1bJ\x04\x08\x1b\x10\x1cJ\x04\x08\x1c\x10\x1dJ\x04\x08\x1d\x10\x1eJ\x04\x08\x1e\x10\x1fJ\x04\x08\x1f\x10 J\x04\x08 \x10!J\x04\x08!\x10\"J\x04\x08#\x10$R\x05stateR\nmessage_toR\x08page_urlR\npage_titleR\x10\x62rowser_languageR\x0b\x64\x65vice_typeR\x07user_idR\tjoiner_idR\x0bjoiner_typeR\tleaver_idR\x0b\x61ttachmentsR\x04textR\x06\x66ormatR\x0b\x64\x65liveriedsR\x06\x66ieldsR\x08\x63omputed\"\xab\x01\n\x05Reply\x12\x1c\n\x03\x63tx\x18\x03 \x01(\x0b\x32\x0f.common.Context\x12\x10\n\x08\x65vent_id\x18\x01 \x01(\t\x12\r\n\x05state\x18\x02 \x01(\x0c\x12\x0b\n\x03\x65rr\x18\n \x01(\x08\x12\x17\n\x0f\x65rr_description\x18\x0c \x01(\t\x12\x19\n\x08\x65rr_code\x18\r \x01(\x0e\x32\x07.lang.T\x12\x11\n\terr_class\x18\x0f \x01(\x05\x12\x0f\n\x07payload\x18\x07 \x01(\x0c\"?\n\x05\x45rror\x12\r\n\x05\x65rror\x18\x02 \x01(\t\x12\x12\n\nrequest_id\x18\x03 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x04 \x01(\t\"\x9b\x01\n\x0cSubscription\x12\x1c\n\x03\x63tx\x18\x01 \x01(\x0b\x32\x0f.common.Context\x12\r\n\x05topic\x18\x03 \x01(\t\x12\x0e\n\x06sub_id\x18\x07 \x01(\t\x12\x14\n\x0ctarget_topic\x18\n \x01(\t\x12\x12\n\ntarget_key\x18\x0b \x01(\t\x12\x0c\n\x04ttls\x18\x0c \x01(\x03J\x04\x08\r\x10\x0eR\x10target_partition*\x7f\n\x05\x45vent\x12\t\n\x05Send_\x10\x01\x12\x0c\n\x08\x41piReply\x10\x02\x12\x13\n\x0fRawEventCreated\x10\x03\x12\r\n\tSubscribe\x10\x04\x12\x12\n\x0eSubscribeReply\x10\x06\x12\x0f\n\x0bUnsubscribe\x10\x05\x12\x14\n\x10UnsubscribeReply\x10\x07*\xc0\x02\n\x04Type\x12\x18\n\x14\x63onversation_updated\x10\t\x12\x10\n\x0cmessage_sent\x10\n\x12\x17\n\x13\x63onversation_joined\x10\x02\x12\x15\n\x11\x63onversation_left\x10\x04\x12\x17\n\x13\x63onversation_tagged\x10\x06\x12\x19\n\x15\x63onversation_untagged\x10\x07\x12\x18\n\x14\x63hannel_deintegrated\x10\x14\x12\x16\n\x12\x63hannel_integrated\x10\x15\x12\x10\n\x0cmessage_seen\x10\x1e\x12\x11\n\rmessage_acked\x10\x1f\x12\x14\n\x10message_received\x10 \x12\x1e\n\x1a\x63onversation_member_typing\x10!\x12\x1b\n\x17\x63onversation_postbacked\x10\"*\'\n\x0eSubscriberType\x12\x08\n\x04user\x10\x00\x12\x0b\n\x07\x63hannel\x10\x01*\'\n\tSubPrefix\x12\x0b\n\x07Webhook\x10\x00\x12\r\n\tWebsocket\x10\x01*^\n\x06Object\x12\x0b\n\x07\x61\x63\x63ount\x10\x00\x12\x10\n\x0c\x63onversation\x10\x01\x12\x0b\n\x07message\x10\x02\x12\x0f\n\x0bintegration\x10\x03\x12\t\n\x05\x61gent\x10\x04\x12\x0c\n\x08postback\x10\x05')
  ,
  dependencies=[bitbucket_dot_org_dot_subiz_dot_header_dot_common_dot_common__pb2.DESCRIPTOR,bitbucket_dot_org_dot_subiz_dot_header_dot_lang_dot_lang__pb2.DESCRIPTOR,bitbucket_dot_org_dot_subiz_dot_header_dot_account_dot_account__pb2.DESCRIPTOR,bitbucket_dot_org_dot_subiz_dot_header_dot_conversation_dot_conversation__pb2.DESCRIPTOR,])

_EVENT = _descriptor.EnumDescriptor(
  name='Event',
  full_name='event.Event',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Send_', index=0, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ApiReply', index=1, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='RawEventCreated', index=2, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Subscribe', index=3, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SubscribeReply', index=4, number=6,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Unsubscribe', index=5, number=5,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UnsubscribeReply', index=6, number=7,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1590,
  serialized_end=1717,
)
_sym_db.RegisterEnumDescriptor(_EVENT)

Event = enum_type_wrapper.EnumTypeWrapper(_EVENT)
_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='event.Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='conversation_updated', index=0, number=9,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='message_sent', index=1, number=10,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='conversation_joined', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='conversation_left', index=3, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='conversation_tagged', index=4, number=6,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='conversation_untagged', index=5, number=7,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='channel_deintegrated', index=6, number=20,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='channel_integrated', index=7, number=21,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='message_seen', index=8, number=30,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='message_acked', index=9, number=31,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='message_received', index=10, number=32,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='conversation_member_typing', index=11, number=33,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='conversation_postbacked', index=12, number=34,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1720,
  serialized_end=2040,
)
_sym_db.RegisterEnumDescriptor(_TYPE)

Type = enum_type_wrapper.EnumTypeWrapper(_TYPE)
_SUBSCRIBERTYPE = _descriptor.EnumDescriptor(
  name='SubscriberType',
  full_name='event.SubscriberType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='user', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='channel', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=2042,
  serialized_end=2081,
)
_sym_db.RegisterEnumDescriptor(_SUBSCRIBERTYPE)

SubscriberType = enum_type_wrapper.EnumTypeWrapper(_SUBSCRIBERTYPE)
_SUBPREFIX = _descriptor.EnumDescriptor(
  name='SubPrefix',
  full_name='event.SubPrefix',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Webhook', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Websocket', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=2083,
  serialized_end=2122,
)
_sym_db.RegisterEnumDescriptor(_SUBPREFIX)

SubPrefix = enum_type_wrapper.EnumTypeWrapper(_SUBPREFIX)
_OBJECT = _descriptor.EnumDescriptor(
  name='Object',
  full_name='event.Object',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='account', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='conversation', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='message', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='integration', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='agent', index=4, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='postback', index=5, number=5,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=2124,
  serialized_end=2218,
)
_sym_db.RegisterEnumDescriptor(_OBJECT)

Object = enum_type_wrapper.EnumTypeWrapper(_OBJECT)
Send_ = 1
ApiReply = 2
RawEventCreated = 3
Subscribe = 4
SubscribeReply = 6
Unsubscribe = 5
UnsubscribeReply = 7
conversation_updated = 9
message_sent = 10
conversation_joined = 2
conversation_left = 4
conversation_tagged = 6
conversation_untagged = 7
channel_deintegrated = 20
channel_integrated = 21
message_seen = 30
message_acked = 31
message_received = 32
conversation_member_typing = 33
conversation_postbacked = 34
user = 0
channel = 1
Webhook = 0
Websocket = 1
account = 0
conversation = 1
message = 2
integration = 3
agent = 4
postback = 5



_RAWEVENTCREATEDPAYLOAD = _descriptor.Descriptor(
  name='RawEventCreatedPayload',
  full_name='event.RawEventCreatedPayload',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ctx', full_name='event.RawEventCreatedPayload.ctx', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='sub_id', full_name='event.RawEventCreatedPayload.sub_id', index=1,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='topic', full_name='event.RawEventCreatedPayload.topic', index=2,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='payload', full_name='event.RawEventCreatedPayload.payload', index=3,
      number=5, type=9, cpp_type=9, label=1,
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
  serialized_start=230,
  serialized_end=332,
)


_UNSUBSCRIBEMESSAGE = _descriptor.Descriptor(
  name='UnsubscribeMessage',
  full_name='event.UnsubscribeMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='topic', full_name='event.UnsubscribeMessage.topic', index=0,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='sub_id', full_name='event.UnsubscribeMessage.sub_id', index=1,
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
  serialized_start=334,
  serialized_end=385,
)


_RAWEVENTS = _descriptor.Descriptor(
  name='RawEvents',
  full_name='event.RawEvents',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='events', full_name='event.RawEvents.events', index=0,
      number=4, type=11, cpp_type=10, label=3,
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
  serialized_start=387,
  serialized_end=431,
)


_BY = _descriptor.Descriptor(
  name='By',
  full_name='event.By',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='device', full_name='event.By.device', index=0,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='member_id', full_name='event.By.member_id', index=1,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='client_id', full_name='event.By.client_id', index=2,
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
  serialized_start=433,
  serialized_end=507,
)


_RAWEVENT = _descriptor.Descriptor(
  name='RawEvent',
  full_name='event.RawEvent',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ctx', full_name='event.RawEvent.ctx', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='id', full_name='event.RawEvent.id', index=1,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='account_id', full_name='event.RawEvent.account_id', index=2,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='created', full_name='event.RawEvent.created', index=3,
      number=8, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='type', full_name='event.RawEvent.type', index=4,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='topics', full_name='event.RawEvent.topics', index=5,
      number=11, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='by', full_name='event.RawEvent.by', index=6,
      number=50, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='object', full_name='event.RawEvent.object', index=7,
      number=13, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='conversation_id', full_name='event.RawEvent.conversation_id', index=8,
      number=20, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='account', full_name='event.RawEvent.account', index=9,
      number=38, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='agent', full_name='event.RawEvent.agent', index=10,
      number=36, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='message', full_name='event.RawEvent.message', index=11,
      number=37, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='conversation', full_name='event.RawEvent.conversation', index=12,
      number=39, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='postback', full_name='event.RawEvent.postback', index=13,
      number=40, type=11, cpp_type=10, label=1,
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
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='data', full_name='event.RawEvent.data',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=510,
  serialized_end=1191,
)


_REPLY = _descriptor.Descriptor(
  name='Reply',
  full_name='event.Reply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ctx', full_name='event.Reply.ctx', index=0,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='event_id', full_name='event.Reply.event_id', index=1,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='state', full_name='event.Reply.state', index=2,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='err', full_name='event.Reply.err', index=3,
      number=10, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='err_description', full_name='event.Reply.err_description', index=4,
      number=12, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='err_code', full_name='event.Reply.err_code', index=5,
      number=13, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='err_class', full_name='event.Reply.err_class', index=6,
      number=15, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='payload', full_name='event.Reply.payload', index=7,
      number=7, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
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
  serialized_start=1194,
  serialized_end=1365,
)


_ERROR = _descriptor.Descriptor(
  name='Error',
  full_name='event.Error',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='error', full_name='event.Error.error', index=0,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='request_id', full_name='event.Error.request_id', index=1,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='description', full_name='event.Error.description', index=2,
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
  serialized_start=1367,
  serialized_end=1430,
)


_SUBSCRIPTION = _descriptor.Descriptor(
  name='Subscription',
  full_name='event.Subscription',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ctx', full_name='event.Subscription.ctx', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='topic', full_name='event.Subscription.topic', index=1,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='sub_id', full_name='event.Subscription.sub_id', index=2,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='target_topic', full_name='event.Subscription.target_topic', index=3,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='target_key', full_name='event.Subscription.target_key', index=4,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ttls', full_name='event.Subscription.ttls', index=5,
      number=12, type=3, cpp_type=2, label=1,
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
  serialized_start=1433,
  serialized_end=1588,
)

_RAWEVENTCREATEDPAYLOAD.fields_by_name['ctx'].message_type = bitbucket_dot_org_dot_subiz_dot_header_dot_common_dot_common__pb2._CONTEXT
_RAWEVENTS.fields_by_name['events'].message_type = _RAWEVENT
_BY.fields_by_name['device'].message_type = bitbucket_dot_org_dot_subiz_dot_header_dot_common_dot_common__pb2._DEVICE
_RAWEVENT.fields_by_name['ctx'].message_type = bitbucket_dot_org_dot_subiz_dot_header_dot_common_dot_common__pb2._CONTEXT
_RAWEVENT.fields_by_name['by'].message_type = _BY
_RAWEVENT.fields_by_name['account'].message_type = bitbucket_dot_org_dot_subiz_dot_header_dot_account_dot_account__pb2._ACCOUNT
_RAWEVENT.fields_by_name['agent'].message_type = bitbucket_dot_org_dot_subiz_dot_header_dot_account_dot_account__pb2._AGENT
_RAWEVENT.fields_by_name['message'].message_type = bitbucket_dot_org_dot_subiz_dot_header_dot_conversation_dot_conversation__pb2._MESSAGE
_RAWEVENT.fields_by_name['conversation'].message_type = bitbucket_dot_org_dot_subiz_dot_header_dot_conversation_dot_conversation__pb2._CONVERSATION
_RAWEVENT.fields_by_name['postback'].message_type = bitbucket_dot_org_dot_subiz_dot_header_dot_conversation_dot_conversation__pb2._POSTBACK
_RAWEVENT.oneofs_by_name['data'].fields.append(
  _RAWEVENT.fields_by_name['account'])
_RAWEVENT.fields_by_name['account'].containing_oneof = _RAWEVENT.oneofs_by_name['data']
_RAWEVENT.oneofs_by_name['data'].fields.append(
  _RAWEVENT.fields_by_name['agent'])
_RAWEVENT.fields_by_name['agent'].containing_oneof = _RAWEVENT.oneofs_by_name['data']
_RAWEVENT.oneofs_by_name['data'].fields.append(
  _RAWEVENT.fields_by_name['message'])
_RAWEVENT.fields_by_name['message'].containing_oneof = _RAWEVENT.oneofs_by_name['data']
_RAWEVENT.oneofs_by_name['data'].fields.append(
  _RAWEVENT.fields_by_name['conversation'])
_RAWEVENT.fields_by_name['conversation'].containing_oneof = _RAWEVENT.oneofs_by_name['data']
_RAWEVENT.oneofs_by_name['data'].fields.append(
  _RAWEVENT.fields_by_name['postback'])
_RAWEVENT.fields_by_name['postback'].containing_oneof = _RAWEVENT.oneofs_by_name['data']
_REPLY.fields_by_name['ctx'].message_type = bitbucket_dot_org_dot_subiz_dot_header_dot_common_dot_common__pb2._CONTEXT
_REPLY.fields_by_name['err_code'].enum_type = bitbucket_dot_org_dot_subiz_dot_header_dot_lang_dot_lang__pb2._T
_SUBSCRIPTION.fields_by_name['ctx'].message_type = bitbucket_dot_org_dot_subiz_dot_header_dot_common_dot_common__pb2._CONTEXT
DESCRIPTOR.message_types_by_name['RawEventCreatedPayload'] = _RAWEVENTCREATEDPAYLOAD
DESCRIPTOR.message_types_by_name['UnsubscribeMessage'] = _UNSUBSCRIBEMESSAGE
DESCRIPTOR.message_types_by_name['RawEvents'] = _RAWEVENTS
DESCRIPTOR.message_types_by_name['By'] = _BY
DESCRIPTOR.message_types_by_name['RawEvent'] = _RAWEVENT
DESCRIPTOR.message_types_by_name['Reply'] = _REPLY
DESCRIPTOR.message_types_by_name['Error'] = _ERROR
DESCRIPTOR.message_types_by_name['Subscription'] = _SUBSCRIPTION
DESCRIPTOR.enum_types_by_name['Event'] = _EVENT
DESCRIPTOR.enum_types_by_name['Type'] = _TYPE
DESCRIPTOR.enum_types_by_name['SubscriberType'] = _SUBSCRIBERTYPE
DESCRIPTOR.enum_types_by_name['SubPrefix'] = _SUBPREFIX
DESCRIPTOR.enum_types_by_name['Object'] = _OBJECT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

RawEventCreatedPayload = _reflection.GeneratedProtocolMessageType('RawEventCreatedPayload', (_message.Message,), dict(
  DESCRIPTOR = _RAWEVENTCREATEDPAYLOAD,
  __module__ = 'event.event_pb2'
  # @@protoc_insertion_point(class_scope:event.RawEventCreatedPayload)
  ))
_sym_db.RegisterMessage(RawEventCreatedPayload)

UnsubscribeMessage = _reflection.GeneratedProtocolMessageType('UnsubscribeMessage', (_message.Message,), dict(
  DESCRIPTOR = _UNSUBSCRIBEMESSAGE,
  __module__ = 'event.event_pb2'
  # @@protoc_insertion_point(class_scope:event.UnsubscribeMessage)
  ))
_sym_db.RegisterMessage(UnsubscribeMessage)

RawEvents = _reflection.GeneratedProtocolMessageType('RawEvents', (_message.Message,), dict(
  DESCRIPTOR = _RAWEVENTS,
  __module__ = 'event.event_pb2'
  # @@protoc_insertion_point(class_scope:event.RawEvents)
  ))
_sym_db.RegisterMessage(RawEvents)

By = _reflection.GeneratedProtocolMessageType('By', (_message.Message,), dict(
  DESCRIPTOR = _BY,
  __module__ = 'event.event_pb2'
  # @@protoc_insertion_point(class_scope:event.By)
  ))
_sym_db.RegisterMessage(By)

RawEvent = _reflection.GeneratedProtocolMessageType('RawEvent', (_message.Message,), dict(
  DESCRIPTOR = _RAWEVENT,
  __module__ = 'event.event_pb2'
  # @@protoc_insertion_point(class_scope:event.RawEvent)
  ))
_sym_db.RegisterMessage(RawEvent)

Reply = _reflection.GeneratedProtocolMessageType('Reply', (_message.Message,), dict(
  DESCRIPTOR = _REPLY,
  __module__ = 'event.event_pb2'
  # @@protoc_insertion_point(class_scope:event.Reply)
  ))
_sym_db.RegisterMessage(Reply)

Error = _reflection.GeneratedProtocolMessageType('Error', (_message.Message,), dict(
  DESCRIPTOR = _ERROR,
  __module__ = 'event.event_pb2'
  # @@protoc_insertion_point(class_scope:event.Error)
  ))
_sym_db.RegisterMessage(Error)

Subscription = _reflection.GeneratedProtocolMessageType('Subscription', (_message.Message,), dict(
  DESCRIPTOR = _SUBSCRIPTION,
  __module__ = 'event.event_pb2'
  # @@protoc_insertion_point(class_scope:event.Subscription)
  ))
_sym_db.RegisterMessage(Subscription)


# @@protoc_insertion_point(module_scope)
