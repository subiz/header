# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: auth/auth.proto

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
  name='auth/auth.proto',
  package='auth',
  syntax='proto3',
  serialized_pb=_b('\n\x0f\x61uth/auth.proto\x12\x04\x61uth\"\x07\n\x05\x45mpty\"\x95\x01\n\nCredential\x12\x11\n\tAccountId\x18\x01 \x01(\t\x12\x0e\n\x06UserId\x18\x03 \x01(\t\x12\x1c\n\x06Method\x18\x05 \x01(\x0b\x32\x0c.auth.Method\x12\x10\n\x08\x43lientId\x18\x07 \x01(\t\x12$\n\nClientType\x18\x08 \x01(\x0e\x32\x10.auth.ClientType\x12\x0e\n\x06\x41PIKey\x18\t \x01(\t\"t\n\x05Scope\x12\n\n\x02Id\x18\x01 \x01(\t\x12\x0c\n\x04Name\x18\x02 \x01(\t\x12\x0f\n\x07LogoUrl\x18\x03 \x01(\t\x12\r\n\x05Title\x18\x04 \x01(\t\x12\x13\n\x0b\x44\x65scription\x18\x05 \x01(\t\x12\x1c\n\x06Method\x18\x06 \x01(\x0b\x32\x0c.auth.Method\"\x9e\x08\n\x06Method\x12\x15\n\rUpdateTrigger\x18\x32 \x01(\x08\x12\x15\n\rDeleteTrigger\x18\x33 \x01(\x08\x12\x15\n\rCreateTrigger\x18\x34 \x01(\x08\x12\x13\n\x0bReadTrigger\x18\x35 \x01(\x08\x12\x13\n\x0bListTrigger\x18\x36 \x01(\x08\x12\x16\n\x0eReadAllTrigger\x18\x37 \x01(\x08\x12\x16\n\x0eListAllTrigger\x18\x38 \x01(\x08\x12\x18\n\x10UpdateAllTrigger\x18\x39 \x01(\x08\x12\x18\n\x10\x44\x65leteAllTrigger\x18: \x01(\x08\x12\x18\n\x10\x43reateAllTrigger\x18; \x01(\x08\x12\x1d\n\x15ReadAgentNotification\x18i \x01(\x08\x12\x1d\n\x15SeenAgentNotification\x18j \x01(\x08\x12\x0c\n\x04Ping\x18k \x01(\x08\x12\x17\n\x0fUpdatePasswords\x18m \x01(\x08\x12\x16\n\x0eUpdatePassword\x18l \x01(\x08\x12\x14\n\x0cInviteAgents\x18\x46 \x01(\x08\x12\x13\n\x0bUpdateAgent\x18H \x01(\x08\x12\x14\n\x0cUpdateAgents\x18I \x01(\x08\x12\x11\n\tReadAgent\x18L \x01(\x08\x12\x12\n\nReadAgents\x18M \x01(\x08\x12\x15\n\rResetPassword\x18N \x01(\x08\x12\x1e\n\x16UpdateAgentsPermission\x18Q \x01(\x08\x12\x19\n\x11UpdateAgentsState\x18V \x01(\x08\x12\x16\n\x0e\x43onfirmAccount\x18x \x01(\x08\x12\x19\n\x11\x43reateAgentGroups\x18| \x01(\x08\x12\x19\n\x11\x44\x65leteAgentGroups\x18} \x01(\x08\x12\x17\n\x0fReadAgentGroups\x18~ \x01(\x08\x12\x19\n\x11UpdateAgentGroups\x18\x7f \x01(\x08\x12\x12\n\nUpdatePlan\x18{ \x01(\x08\x12\x16\n\rCreateAccount\x18\x81\x01 \x01(\x08\x12\x16\n\rDeleteAccount\x18\x82\x01 \x01(\x08\x12\x15\n\rUpdateAccount\x18z \x01(\x08\x12\x14\n\x0bReadAccount\x18\x83\x01 \x01(\x08\x12\x13\n\nReadClient\x18\x97\x01 \x01(\x08\x12\x13\n\nListClient\x18\x98\x01 \x01(\x08\x12\x19\n\x10UpdateClientInfo\x18\x99\x01 \x01(\x08\x12\x16\n\rPublishClient\x18\x9a\x01 \x01(\x08\x12\x1f\n\x16RegenerateClientSecret\x18\x9b\x01 \x01(\x08\x12\x15\n\x0cVerifyClient\x18\x9c\x01 \x01(\x08\x12\x15\n\x0c\x44\x65leteClient\x18\x9d\x01 \x01(\x08\x12\x15\n\x0c\x43reateClient\x18\x9e\x01 \x01(\x08\x12\x11\n\x08ReadRule\x18\xb4\x01 \x01(\x08\x12\x12\n\tWriteRule\x18\xb5\x01 \x01(\x08\x12\x13\n\nDeleteRule\x18\xb6\x01 \x01(\x08\"\x15\n\x07ScopeId\x12\n\n\x02Id\x18\x01 \x01(\t\"\x14\n\x06UserId\x12\n\n\x02Id\x18\x01 \x01(\t\"-\n\x0eScopesResponse\x12\x1b\n\x06Scopes\x18\x01 \x03(\x0b\x32\x0b.auth.Scope\"8\n\x08UserAuth\x12\x0e\n\x06UserId\x18\x01 \x01(\t\x12\x1c\n\x06Method\x18\x02 \x01(\x0b\x32\x0c.auth.Method*J\n\nClientType\x12\x0b\n\x07Unknown\x10\x00\x12\x0b\n\x07\x43hannel\x10\x01\x12\n\n\x06Widget\x10\x02\x12\r\n\tDashboard\x10\x03\x12\x07\n\x03\x41pp\x10\x04\x32^\n\x04\x41uth\x12)\n\tReadScope\x12\r.auth.ScopeId\x1a\x0b.auth.Scope\"\x00\x12+\n\x04List\x12\x0b.auth.Empty\x1a\x14.auth.ScopesResponse\"\x00\x62\x06proto3')
)
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

_CLIENTTYPE = _descriptor.EnumDescriptor(
  name='ClientType',
  full_name='auth.ClientType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Unknown', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Channel', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Widget', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Dashboard', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='App', index=4, number=4,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1511,
  serialized_end=1585,
)
_sym_db.RegisterEnumDescriptor(_CLIENTTYPE)

ClientType = enum_type_wrapper.EnumTypeWrapper(_CLIENTTYPE)
Unknown = 0
Channel = 1
Widget = 2
Dashboard = 3
App = 4



_EMPTY = _descriptor.Descriptor(
  name='Empty',
  full_name='auth.Empty',
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
  serialized_start=25,
  serialized_end=32,
)


_CREDENTIAL = _descriptor.Descriptor(
  name='Credential',
  full_name='auth.Credential',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='AccountId', full_name='auth.Credential.AccountId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='UserId', full_name='auth.Credential.UserId', index=1,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Method', full_name='auth.Credential.Method', index=2,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ClientId', full_name='auth.Credential.ClientId', index=3,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ClientType', full_name='auth.Credential.ClientType', index=4,
      number=8, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='APIKey', full_name='auth.Credential.APIKey', index=5,
      number=9, type=9, cpp_type=9, label=1,
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
  serialized_start=35,
  serialized_end=184,
)


_SCOPE = _descriptor.Descriptor(
  name='Scope',
  full_name='auth.Scope',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Id', full_name='auth.Scope.Id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Name', full_name='auth.Scope.Name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='LogoUrl', full_name='auth.Scope.LogoUrl', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Title', full_name='auth.Scope.Title', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Description', full_name='auth.Scope.Description', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Method', full_name='auth.Scope.Method', index=5,
      number=6, type=11, cpp_type=10, label=1,
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
  serialized_start=186,
  serialized_end=302,
)


_METHOD = _descriptor.Descriptor(
  name='Method',
  full_name='auth.Method',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='UpdateTrigger', full_name='auth.Method.UpdateTrigger', index=0,
      number=50, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='DeleteTrigger', full_name='auth.Method.DeleteTrigger', index=1,
      number=51, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='CreateTrigger', full_name='auth.Method.CreateTrigger', index=2,
      number=52, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ReadTrigger', full_name='auth.Method.ReadTrigger', index=3,
      number=53, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ListTrigger', full_name='auth.Method.ListTrigger', index=4,
      number=54, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ReadAllTrigger', full_name='auth.Method.ReadAllTrigger', index=5,
      number=55, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ListAllTrigger', full_name='auth.Method.ListAllTrigger', index=6,
      number=56, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='UpdateAllTrigger', full_name='auth.Method.UpdateAllTrigger', index=7,
      number=57, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='DeleteAllTrigger', full_name='auth.Method.DeleteAllTrigger', index=8,
      number=58, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='CreateAllTrigger', full_name='auth.Method.CreateAllTrigger', index=9,
      number=59, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ReadAgentNotification', full_name='auth.Method.ReadAgentNotification', index=10,
      number=105, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='SeenAgentNotification', full_name='auth.Method.SeenAgentNotification', index=11,
      number=106, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Ping', full_name='auth.Method.Ping', index=12,
      number=107, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='UpdatePasswords', full_name='auth.Method.UpdatePasswords', index=13,
      number=109, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='UpdatePassword', full_name='auth.Method.UpdatePassword', index=14,
      number=108, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='InviteAgents', full_name='auth.Method.InviteAgents', index=15,
      number=70, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='UpdateAgent', full_name='auth.Method.UpdateAgent', index=16,
      number=72, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='UpdateAgents', full_name='auth.Method.UpdateAgents', index=17,
      number=73, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ReadAgent', full_name='auth.Method.ReadAgent', index=18,
      number=76, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ReadAgents', full_name='auth.Method.ReadAgents', index=19,
      number=77, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ResetPassword', full_name='auth.Method.ResetPassword', index=20,
      number=78, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='UpdateAgentsPermission', full_name='auth.Method.UpdateAgentsPermission', index=21,
      number=81, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='UpdateAgentsState', full_name='auth.Method.UpdateAgentsState', index=22,
      number=86, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ConfirmAccount', full_name='auth.Method.ConfirmAccount', index=23,
      number=120, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='CreateAgentGroups', full_name='auth.Method.CreateAgentGroups', index=24,
      number=124, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='DeleteAgentGroups', full_name='auth.Method.DeleteAgentGroups', index=25,
      number=125, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ReadAgentGroups', full_name='auth.Method.ReadAgentGroups', index=26,
      number=126, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='UpdateAgentGroups', full_name='auth.Method.UpdateAgentGroups', index=27,
      number=127, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='UpdatePlan', full_name='auth.Method.UpdatePlan', index=28,
      number=123, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='CreateAccount', full_name='auth.Method.CreateAccount', index=29,
      number=129, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='DeleteAccount', full_name='auth.Method.DeleteAccount', index=30,
      number=130, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='UpdateAccount', full_name='auth.Method.UpdateAccount', index=31,
      number=122, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ReadAccount', full_name='auth.Method.ReadAccount', index=32,
      number=131, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ReadClient', full_name='auth.Method.ReadClient', index=33,
      number=151, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ListClient', full_name='auth.Method.ListClient', index=34,
      number=152, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='UpdateClientInfo', full_name='auth.Method.UpdateClientInfo', index=35,
      number=153, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='PublishClient', full_name='auth.Method.PublishClient', index=36,
      number=154, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='RegenerateClientSecret', full_name='auth.Method.RegenerateClientSecret', index=37,
      number=155, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='VerifyClient', full_name='auth.Method.VerifyClient', index=38,
      number=156, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='DeleteClient', full_name='auth.Method.DeleteClient', index=39,
      number=157, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='CreateClient', full_name='auth.Method.CreateClient', index=40,
      number=158, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='ReadRule', full_name='auth.Method.ReadRule', index=41,
      number=180, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='WriteRule', full_name='auth.Method.WriteRule', index=42,
      number=181, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='DeleteRule', full_name='auth.Method.DeleteRule', index=43,
      number=182, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=305,
  serialized_end=1359,
)


_SCOPEID = _descriptor.Descriptor(
  name='ScopeId',
  full_name='auth.ScopeId',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Id', full_name='auth.ScopeId.Id', index=0,
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
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1361,
  serialized_end=1382,
)


_USERID = _descriptor.Descriptor(
  name='UserId',
  full_name='auth.UserId',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Id', full_name='auth.UserId.Id', index=0,
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
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1384,
  serialized_end=1404,
)


_SCOPESRESPONSE = _descriptor.Descriptor(
  name='ScopesResponse',
  full_name='auth.ScopesResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Scopes', full_name='auth.ScopesResponse.Scopes', index=0,
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
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1406,
  serialized_end=1451,
)


_USERAUTH = _descriptor.Descriptor(
  name='UserAuth',
  full_name='auth.UserAuth',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='UserId', full_name='auth.UserAuth.UserId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Method', full_name='auth.UserAuth.Method', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_start=1453,
  serialized_end=1509,
)

_CREDENTIAL.fields_by_name['Method'].message_type = _METHOD
_CREDENTIAL.fields_by_name['ClientType'].enum_type = _CLIENTTYPE
_SCOPE.fields_by_name['Method'].message_type = _METHOD
_SCOPESRESPONSE.fields_by_name['Scopes'].message_type = _SCOPE
_USERAUTH.fields_by_name['Method'].message_type = _METHOD
DESCRIPTOR.message_types_by_name['Empty'] = _EMPTY
DESCRIPTOR.message_types_by_name['Credential'] = _CREDENTIAL
DESCRIPTOR.message_types_by_name['Scope'] = _SCOPE
DESCRIPTOR.message_types_by_name['Method'] = _METHOD
DESCRIPTOR.message_types_by_name['ScopeId'] = _SCOPEID
DESCRIPTOR.message_types_by_name['UserId'] = _USERID
DESCRIPTOR.message_types_by_name['ScopesResponse'] = _SCOPESRESPONSE
DESCRIPTOR.message_types_by_name['UserAuth'] = _USERAUTH
DESCRIPTOR.enum_types_by_name['ClientType'] = _CLIENTTYPE

Empty = _reflection.GeneratedProtocolMessageType('Empty', (_message.Message,), dict(
  DESCRIPTOR = _EMPTY,
  __module__ = 'auth.auth_pb2'
  # @@protoc_insertion_point(class_scope:auth.Empty)
  ))
_sym_db.RegisterMessage(Empty)

Credential = _reflection.GeneratedProtocolMessageType('Credential', (_message.Message,), dict(
  DESCRIPTOR = _CREDENTIAL,
  __module__ = 'auth.auth_pb2'
  # @@protoc_insertion_point(class_scope:auth.Credential)
  ))
_sym_db.RegisterMessage(Credential)

Scope = _reflection.GeneratedProtocolMessageType('Scope', (_message.Message,), dict(
  DESCRIPTOR = _SCOPE,
  __module__ = 'auth.auth_pb2'
  # @@protoc_insertion_point(class_scope:auth.Scope)
  ))
_sym_db.RegisterMessage(Scope)

Method = _reflection.GeneratedProtocolMessageType('Method', (_message.Message,), dict(
  DESCRIPTOR = _METHOD,
  __module__ = 'auth.auth_pb2'
  # @@protoc_insertion_point(class_scope:auth.Method)
  ))
_sym_db.RegisterMessage(Method)

ScopeId = _reflection.GeneratedProtocolMessageType('ScopeId', (_message.Message,), dict(
  DESCRIPTOR = _SCOPEID,
  __module__ = 'auth.auth_pb2'
  # @@protoc_insertion_point(class_scope:auth.ScopeId)
  ))
_sym_db.RegisterMessage(ScopeId)

UserId = _reflection.GeneratedProtocolMessageType('UserId', (_message.Message,), dict(
  DESCRIPTOR = _USERID,
  __module__ = 'auth.auth_pb2'
  # @@protoc_insertion_point(class_scope:auth.UserId)
  ))
_sym_db.RegisterMessage(UserId)

ScopesResponse = _reflection.GeneratedProtocolMessageType('ScopesResponse', (_message.Message,), dict(
  DESCRIPTOR = _SCOPESRESPONSE,
  __module__ = 'auth.auth_pb2'
  # @@protoc_insertion_point(class_scope:auth.ScopesResponse)
  ))
_sym_db.RegisterMessage(ScopesResponse)

UserAuth = _reflection.GeneratedProtocolMessageType('UserAuth', (_message.Message,), dict(
  DESCRIPTOR = _USERAUTH,
  __module__ = 'auth.auth_pb2'
  # @@protoc_insertion_point(class_scope:auth.UserAuth)
  ))
_sym_db.RegisterMessage(UserAuth)


# @@protoc_insertion_point(module_scope)
