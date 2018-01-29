// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lang/lang.proto

/*
Package lang is a generated protocol buffer package.

It is generated from these files:
	lang/lang.proto

It has these top-level messages:
*/
package lang

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type T int32

const (
	T_undefined                         T = 0
	T_user_has_already_in_conversation  T = 1
	T_conversation_closed               T = 2
	T_invalid_invite                    T = 3
	T_invalid_agent                     T = 4
	T_member_is_not_in_conversation     T = 5
	T_conversation_not_found            T = 6
	T_internal_error                    T = 30
	T_invalid_input                     T = 22
	T_invalid_form                      T = 20
	T_access_token_expired              T = 21
	T_credential_not_set                T = 7
	T_wrong_account_in_credential       T = 8
	T_access_deny                       T = 9
	T_wrong_user_in_credential          T = 10
	T_unable_to_send_message            T = 11
	T_topic_is_empty                    T = 12
	T_invalid_credential                T = 13
	T_invalid_left                      T = 14
	T_invalid_json                      T = 15
	T_invalid_protobuf                  T = 16
	T_invalid_password                  T = 17
	T_wrong_password                    T = 18
	T_invalid_agent_state               T = 19
	T_unable_to_lock                    T = 40
	T_empty                             T = 41
	T_wrong_type                        T = 42
	T_invalid_kafka_topic               T = 43
	T_database_error                    T = 44
	T_timeout                           T = 45
	T_websocket_error                   T = 46
	T_kafka_error                       T = 47
	T_invalid_token                     T = 48
	T_account_not_found                 T = 49
	T_agent_not_found                   T = 50
	T_invalid_email                     T = 60
	T_plan_not_found                    T = 61
	T_agent_group_not_found             T = 62
	T_empty_client_type                 T = 63
	T_empty_url                         T = 64
	T_empty_name                        T = 65
	T_client_not_found                  T = 66
	T_empty_account                     T = 70
	T_invalid_conversation_state        T = 71
	T_invalid_message_id                T = 80
	T_invalid_mask                      T = 81
	T_randomize_error                   T = 82
	T_duplicated_message_received_error T = 83
	T_network_error                     T = 84
	T_rsa_key_not_found                 T = 85
	T_jwt_sign_error                    T = 86
	T_env_config_error                  T = 87
	T_scrypt_error                      T = 90
	T_challenge_not_found               T = 91
	T_wrong_answer                      T = 92
	T_server_listen_error               T = 93
	T_scrypt_file_not_found             T = 94
	T_invalid_topic                     T = 95
	T_file_not_found                    T = 99
	T_user_not_found                    T = 100
	T_empty_md5                         T = 101
	T_base_convert_error                T = 102
	T_s3_error                          T = 103
	T_aws_credential_error              T = 104
	T_aws_send_error                    T = 105
	T_elasticsearch_error               T = 200
	T_invalid_template                  T = 203
	T_sendgrid_send_error               T = 204
	T_whitelist_domain_not_found        T = 205
	T_blacklist_ip_not_found            T = 206
	T_blocked_user_not_found            T = 207
	T_invalid_content_type              T = 210
	T_parse_file_error                  T = 211
	T_invalid_integration_id            T = 220
	T_invalid_integration               T = 221
	T_webhook_not_found                 T = 222
	T_tempfile_error                    T = 223
	T_write_file_error                  T = 224
	T_close_file_error                  T = 225
	T_execute_shell_error               T = 226
	T_invalid_css                       T = 227
	T_invalid_hmac                      T = 228
	T_consul_error                      T = 230
	T_maxminddb_err                     T = 231
	T_invalid_condition_key             T = 232
	T_invalid_po_file                   T = 233
	T_integration_not_found             T = 234
	T_webhook_is_disabled               T = 235
)

var T_name = map[int32]string{
	0:   "undefined",
	1:   "user_has_already_in_conversation",
	2:   "conversation_closed",
	3:   "invalid_invite",
	4:   "invalid_agent",
	5:   "member_is_not_in_conversation",
	6:   "conversation_not_found",
	30:  "internal_error",
	22:  "invalid_input",
	20:  "invalid_form",
	21:  "access_token_expired",
	7:   "credential_not_set",
	8:   "wrong_account_in_credential",
	9:   "access_deny",
	10:  "wrong_user_in_credential",
	11:  "unable_to_send_message",
	12:  "topic_is_empty",
	13:  "invalid_credential",
	14:  "invalid_left",
	15:  "invalid_json",
	16:  "invalid_protobuf",
	17:  "invalid_password",
	18:  "wrong_password",
	19:  "invalid_agent_state",
	40:  "unable_to_lock",
	41:  "empty",
	42:  "wrong_type",
	43:  "invalid_kafka_topic",
	44:  "database_error",
	45:  "timeout",
	46:  "websocket_error",
	47:  "kafka_error",
	48:  "invalid_token",
	49:  "account_not_found",
	50:  "agent_not_found",
	60:  "invalid_email",
	61:  "plan_not_found",
	62:  "agent_group_not_found",
	63:  "empty_client_type",
	64:  "empty_url",
	65:  "empty_name",
	66:  "client_not_found",
	70:  "empty_account",
	71:  "invalid_conversation_state",
	80:  "invalid_message_id",
	81:  "invalid_mask",
	82:  "randomize_error",
	83:  "duplicated_message_received_error",
	84:  "network_error",
	85:  "rsa_key_not_found",
	86:  "jwt_sign_error",
	87:  "env_config_error",
	90:  "scrypt_error",
	91:  "challenge_not_found",
	92:  "wrong_answer",
	93:  "server_listen_error",
	94:  "scrypt_file_not_found",
	95:  "invalid_topic",
	99:  "file_not_found",
	100: "user_not_found",
	101: "empty_md5",
	102: "base_convert_error",
	103: "s3_error",
	104: "aws_credential_error",
	105: "aws_send_error",
	200: "elasticsearch_error",
	203: "invalid_template",
	204: "sendgrid_send_error",
	205: "whitelist_domain_not_found",
	206: "blacklist_ip_not_found",
	207: "blocked_user_not_found",
	210: "invalid_content_type",
	211: "parse_file_error",
	220: "invalid_integration_id",
	221: "invalid_integration",
	222: "webhook_not_found",
	223: "tempfile_error",
	224: "write_file_error",
	225: "close_file_error",
	226: "execute_shell_error",
	227: "invalid_css",
	228: "invalid_hmac",
	230: "consul_error",
	231: "maxminddb_err",
	232: "invalid_condition_key",
	233: "invalid_po_file",
	234: "integration_not_found",
	235: "webhook_is_disabled",
}
var T_value = map[string]int32{
	"undefined":                         0,
	"user_has_already_in_conversation":  1,
	"conversation_closed":               2,
	"invalid_invite":                    3,
	"invalid_agent":                     4,
	"member_is_not_in_conversation":     5,
	"conversation_not_found":            6,
	"internal_error":                    30,
	"invalid_input":                     22,
	"invalid_form":                      20,
	"access_token_expired":              21,
	"credential_not_set":                7,
	"wrong_account_in_credential":       8,
	"access_deny":                       9,
	"wrong_user_in_credential":          10,
	"unable_to_send_message":            11,
	"topic_is_empty":                    12,
	"invalid_credential":                13,
	"invalid_left":                      14,
	"invalid_json":                      15,
	"invalid_protobuf":                  16,
	"invalid_password":                  17,
	"wrong_password":                    18,
	"invalid_agent_state":               19,
	"unable_to_lock":                    40,
	"empty":                             41,
	"wrong_type":                        42,
	"invalid_kafka_topic":               43,
	"database_error":                    44,
	"timeout":                           45,
	"websocket_error":                   46,
	"kafka_error":                       47,
	"invalid_token":                     48,
	"account_not_found":                 49,
	"agent_not_found":                   50,
	"invalid_email":                     60,
	"plan_not_found":                    61,
	"agent_group_not_found":             62,
	"empty_client_type":                 63,
	"empty_url":                         64,
	"empty_name":                        65,
	"client_not_found":                  66,
	"empty_account":                     70,
	"invalid_conversation_state":        71,
	"invalid_message_id":                80,
	"invalid_mask":                      81,
	"randomize_error":                   82,
	"duplicated_message_received_error": 83,
	"network_error":                     84,
	"rsa_key_not_found":                 85,
	"jwt_sign_error":                    86,
	"env_config_error":                  87,
	"scrypt_error":                      90,
	"challenge_not_found":               91,
	"wrong_answer":                      92,
	"server_listen_error":               93,
	"scrypt_file_not_found":             94,
	"invalid_topic":                     95,
	"file_not_found":                    99,
	"user_not_found":                    100,
	"empty_md5":                         101,
	"base_convert_error":                102,
	"s3_error":                          103,
	"aws_credential_error":              104,
	"aws_send_error":                    105,
	"elasticsearch_error":               200,
	"invalid_template":                  203,
	"sendgrid_send_error":               204,
	"whitelist_domain_not_found":        205,
	"blacklist_ip_not_found":            206,
	"blocked_user_not_found":            207,
	"invalid_content_type":              210,
	"parse_file_error":                  211,
	"invalid_integration_id":            220,
	"invalid_integration":               221,
	"webhook_not_found":                 222,
	"tempfile_error":                    223,
	"write_file_error":                  224,
	"close_file_error":                  225,
	"execute_shell_error":               226,
	"invalid_css":                       227,
	"invalid_hmac":                      228,
	"consul_error":                      230,
	"maxminddb_err":                     231,
	"invalid_condition_key":             232,
	"invalid_po_file":                   233,
	"integration_not_found":             234,
	"webhook_is_disabled":               235,
}

func (x T) String() string {
	return proto.EnumName(T_name, int32(x))
}
func (T) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type L int32

const (
	L_en L = 0
	L_vn L = 1
	L_cn L = 3
	L_fr L = 5
)

var L_name = map[int32]string{
	0: "en",
	1: "vn",
	3: "cn",
	5: "fr",
}
var L_value = map[string]int32{
	"en": 0,
	"vn": 1,
	"cn": 3,
	"fr": 5,
}

func (x L) String() string {
	return proto.EnumName(L_name, int32(x))
}
func (L) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterEnum("lang.T", T_name, T_value)
	proto.RegisterEnum("lang.L", L_name, L_value)
}

func init() { proto.RegisterFile("lang/lang.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1027 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x55, 0x59, 0x93, 0xdc, 0x34,
	0x10, 0x8e, 0x73, 0x6c, 0x12, 0xed, 0x31, 0xbd, 0xda, 0x83, 0xcd, 0x06, 0x12, 0x52, 0x40, 0x15,
	0x2c, 0x90, 0x00, 0x29, 0xde, 0xb8, 0x1f, 0xe0, 0x85, 0x07, 0x8e, 0x00, 0x55, 0x5c, 0x2a, 0x8d,
	0xd5, 0x33, 0xa3, 0xd8, 0x96, 0x5c, 0x92, 0x3c, 0xb3, 0xc3, 0x1f, 0x84, 0x2a, 0xae, 0x2a, 0x78,
	0xe5, 0x86, 0xe2, 0xfc, 0x13, 0x54, 0xcb, 0xf6, 0x58, 0x13, 0x78, 0x19, 0x8f, 0x3f, 0xb5, 0xba,
	0xbf, 0xaf, 0xfb, 0x93, 0xcc, 0x46, 0xa5, 0x34, 0xd3, 0x5b, 0xf4, 0x73, 0xb3, 0x76, 0x36, 0x58,
	0x7e, 0x9e, 0xfe, 0x9f, 0x7c, 0x3a, 0x62, 0xd9, 0x1d, 0xbe, 0xcd, 0x2e, 0x37, 0x46, 0xe1, 0x44,
	0x1b, 0x54, 0x70, 0x86, 0x3f, 0xcc, 0x1e, 0x6c, 0x3c, 0x3a, 0x31, 0x93, 0x5e, 0xc8, 0xd2, 0xa1,
	0x54, 0x4b, 0xa1, 0x8d, 0xc8, 0xad, 0x99, 0xa3, 0xf3, 0x32, 0x68, 0x6b, 0x20, 0xe3, 0xf7, 0xb1,
	0xbd, 0x14, 0x11, 0x79, 0x69, 0x3d, 0x2a, 0x38, 0xcb, 0x39, 0xdb, 0xd1, 0x66, 0x2e, 0x4b, 0xad,
	0x84, 0x36, 0x73, 0x1d, 0x10, 0xce, 0xf1, 0x5d, 0xb6, 0xdd, 0x63, 0x72, 0x8a, 0x26, 0xc0, 0x79,
	0x7e, 0x83, 0x3d, 0x50, 0x61, 0x35, 0x46, 0x27, 0xb4, 0x17, 0xc6, 0x86, 0xff, 0x94, 0xb8, 0xc0,
	0x8f, 0xd9, 0xe1, 0x5a, 0x09, 0x8a, 0x9a, 0xd8, 0xc6, 0x28, 0xd8, 0x68, 0xab, 0x04, 0x74, 0x46,
	0x96, 0x02, 0x9d, 0xb3, 0x0e, 0xae, 0xa5, 0x55, 0xb4, 0xa9, 0x9b, 0x00, 0x87, 0x1c, 0xd8, 0x56,
	0x0f, 0x4d, 0xac, 0xab, 0x60, 0x9f, 0x1f, 0xb1, 0x7d, 0x99, 0xe7, 0xe8, 0xbd, 0x08, 0xb6, 0x40,
	0x23, 0xf0, 0xb4, 0xd6, 0x0e, 0x15, 0x1c, 0xf0, 0x43, 0xc6, 0x73, 0x87, 0x0a, 0x4d, 0xd0, 0xb2,
	0x8c, 0xc5, 0x3c, 0x06, 0xb8, 0xc8, 0xaf, 0xb3, 0xab, 0x0b, 0x67, 0xcd, 0x54, 0xc8, 0x3c, 0xb7,
	0x8d, 0x69, 0x99, 0xae, 0x02, 0xe1, 0x12, 0x1f, 0xb1, 0xcd, 0x2e, 0xa5, 0x42, 0xb3, 0x84, 0xcb,
	0xfc, 0x7e, 0x76, 0xd4, 0xee, 0x88, 0x7d, 0x5c, 0x0f, 0x67, 0x24, 0xab, 0x31, 0x72, 0x5c, 0xa2,
	0x08, 0x56, 0x78, 0x34, 0x4a, 0x54, 0xe8, 0xbd, 0x9c, 0x22, 0x6c, 0x92, 0xac, 0x60, 0x6b, 0x9d,
	0x53, 0x53, 0xb0, 0xaa, 0xc3, 0x12, 0xb6, 0x88, 0x57, 0xaf, 0x21, 0xc9, 0xb3, 0x9d, 0x6a, 0x2b,
	0x71, 0x12, 0x60, 0x27, 0x45, 0xee, 0x7a, 0x6b, 0x60, 0xc4, 0xf7, 0x19, 0xf4, 0x48, 0x9c, 0xfb,
	0xb8, 0x99, 0x00, 0xac, 0xa1, 0xd2, 0xfb, 0x85, 0x75, 0x0a, 0x76, 0xa9, 0x76, 0xcb, 0x7a, 0x85,
	0x71, 0x9a, 0xf2, 0xda, 0xe0, 0x84, 0x0f, 0x32, 0x20, 0xec, 0x51, 0xf0, 0x20, 0xa2, 0xb4, 0x79,
	0x01, 0x8f, 0xf2, 0xcb, 0xec, 0x42, 0xcb, 0xf9, 0x31, 0xbe, 0xc3, 0x58, 0x9b, 0x2b, 0x2c, 0x6b,
	0x84, 0x93, 0x34, 0x4f, 0x21, 0x27, 0x85, 0x14, 0x51, 0x25, 0x3c, 0x4e, 0x79, 0x94, 0x0c, 0x72,
	0x2c, 0x3d, 0x76, 0x73, 0x7c, 0x82, 0x6f, 0xb2, 0x8b, 0x41, 0x57, 0x68, 0x9b, 0x00, 0x4f, 0xf2,
	0x3d, 0x36, 0x5a, 0xe0, 0xd8, 0xdb, 0xbc, 0xc0, 0xd0, 0x45, 0xdc, 0xa4, 0x8e, 0xb7, 0x69, 0x5a,
	0xe0, 0x56, 0x3a, 0xfa, 0x38, 0x56, 0x78, 0x8a, 0x1f, 0xb0, 0xdd, 0x7e, 0x60, 0x83, 0x71, 0x9e,
	0xa6, 0x7c, 0xad, 0x92, 0x01, 0x7c, 0x26, 0xdd, 0x8e, 0x95, 0xd4, 0x25, 0x3c, 0x47, 0xc4, 0xea,
	0x52, 0xa6, 0xa6, 0x7b, 0x9e, 0x5f, 0x61, 0x07, 0xed, 0xde, 0xa9, 0xb3, 0x4d, 0x9d, 0x2c, 0xbd,
	0x40, 0xd5, 0xa2, 0x76, 0x91, 0x97, 0x9a, 0x22, 0xa2, 0xee, 0x17, 0xe9, 0x68, 0xb5, 0x70, 0xe3,
	0x4a, 0x78, 0x89, 0xda, 0xd2, 0xbe, 0x1a, 0x59, 0x21, 0xbc, 0x4c, 0x83, 0xe8, 0xe2, 0x87, 0x5c,
	0xaf, 0x10, 0x9b, 0x36, 0xaa, 0xe3, 0x0f, 0xaf, 0xf2, 0x6b, 0xec, 0x78, 0xe5, 0x81, 0xf4, 0x48,
	0xb4, 0xe3, 0x78, 0x2d, 0xf5, 0x48, 0x67, 0x26, 0xa1, 0x15, 0xbc, 0x91, 0x3a, 0xa2, 0x92, 0xbe,
	0x80, 0x37, 0x49, 0xbf, 0x93, 0x46, 0xd9, 0x4a, 0x7f, 0xd2, 0x77, 0xfc, 0x2d, 0xfe, 0x08, 0xbb,
	0xa1, 0x9a, 0xba, 0xd4, 0xb9, 0x0c, 0x38, 0x64, 0x70, 0x98, 0xa3, 0x9e, 0xa3, 0xea, 0xc2, 0xde,
	0x26, 0x62, 0x06, 0xc3, 0xc2, 0xba, 0xa2, 0x83, 0xee, 0x90, 0x6e, 0xe7, 0xa5, 0x28, 0x70, 0x99,
	0x48, 0x78, 0x87, 0xba, 0x77, 0x77, 0x11, 0x84, 0xd7, 0x53, 0xd3, 0x85, 0xbe, 0x4b, 0x62, 0xd1,
	0xcc, 0x89, 0xff, 0x44, 0x4f, 0x3b, 0xf4, 0x3d, 0x62, 0xe8, 0x73, 0xb7, 0xac, 0xfb, 0xe1, 0xbe,
	0x1f, 0x6f, 0x96, 0x99, 0x2c, 0x4b, 0x34, 0x53, 0x4c, 0x92, 0x7e, 0x40, 0xa1, 0xdd, 0x41, 0x34,
	0x7e, 0x81, 0x0e, 0x3e, 0xa4, 0x50, 0x8f, 0x6e, 0x8e, 0x4e, 0x94, 0xda, 0x07, 0xec, 0x6b, 0x7d,
	0x44, 0x93, 0xea, 0xb2, 0x4e, 0x74, 0x99, 0x66, 0xf9, 0x78, 0xdd, 0x2a, 0x64, 0x42, 0x41, 0x6c,
	0xef, 0x09, 0xcb, 0xa3, 0xc1, 0xe9, 0xf4, 0x0e, 0x98, 0x1a, 0xa6, 0x59, 0xa9, 0x67, 0x01, 0xa9,
	0xe9, 0xd1, 0xb7, 0xed, 0x44, 0x7a, 0x01, 0x13, 0xbe, 0xc5, 0x2e, 0xf9, 0xdb, 0xdd, 0xdb, 0x34,
	0x5e, 0x38, 0x0b, 0x9f, 0x1c, 0xdd, 0x6e, 0x65, 0x46, 0x25, 0x68, 0x25, 0x5e, 0x01, 0x2d, 0xa6,
	0xf9, 0x11, 0xdb, 0xc3, 0x52, 0xfa, 0xa0, 0x73, 0x8f, 0xd2, 0xe5, 0xb3, 0x6e, 0xe1, 0xb3, 0x8c,
	0x1f, 0x0c, 0x87, 0x36, 0x60, 0x55, 0x97, 0x34, 0xf8, 0xcf, 0x33, 0xda, 0x40, 0x09, 0xa6, 0x4e,
	0xab, 0x34, 0xd3, 0x17, 0x19, 0xbf, 0xce, 0x8e, 0x17, 0x33, 0x1d, 0x90, 0x5a, 0x23, 0x94, 0xad,
	0xa4, 0x4e, 0xdd, 0xfc, 0x65, 0xc6, 0xaf, 0xb2, 0xc3, 0x71, 0x29, 0xf3, 0x22, 0x06, 0xe8, 0xd4,
	0xcf, 0x5f, 0x75, 0x8b, 0x74, 0xea, 0x94, 0xb8, 0xa7, 0x0f, 0x5f, 0x67, 0xfc, 0x0a, 0xdb, 0x4f,
	0xec, 0x18, 0x56, 0x86, 0xff, 0x26, 0xd2, 0xac, 0xa5, 0xf3, 0xd8, 0x36, 0xbe, 0x25, 0xf3, 0x6d,
	0x4c, 0x37, 0xdc, 0xcd, 0x01, 0xa7, 0xae, 0xf5, 0xaf, 0x56, 0xf0, 0x5d, 0xd4, 0xf0, 0x3f, 0x8b,
	0xf0, 0x7d, 0xc6, 0x0f, 0xd9, 0xee, 0x02, 0xc7, 0x33, 0x6b, 0x8b, 0x84, 0xc0, 0x0f, 0x19, 0xdf,
	0x63, 0x3b, 0xd4, 0x84, 0xa4, 0xc6, 0x8f, 0xb1, 0xf4, 0xc2, 0xe9, 0xb0, 0x56, 0xfa, 0xa7, 0x08,
	0xc7, 0x8f, 0x53, 0x0a, 0xff, 0x1c, 0x8b, 0xe2, 0x29, 0xe6, 0x4d, 0x40, 0xe1, 0x67, 0x58, 0xf6,
	0x63, 0xf9, 0x25, 0xe3, 0xc0, 0x36, 0x57, 0xea, 0xbc, 0x87, 0x5f, 0x33, 0xbe, 0x3b, 0x1c, 0xa3,
	0x59, 0x25, 0x73, 0xf8, 0x2d, 0x42, 0xb9, 0x35, 0xbe, 0xe9, 0xf7, 0xfd, 0x9e, 0x71, 0xce, 0xb6,
	0x2b, 0x79, 0x5a, 0x69, 0xa3, 0xd4, 0x98, 0x50, 0xf8, 0x23, 0xe3, 0xc7, 0xec, 0x20, 0xe9, 0x94,
	0xd2, 0x51, 0x75, 0x81, 0x4b, 0xf8, 0x33, 0xe3, 0xfb, 0x6c, 0xb4, 0xba, 0x86, 0x6d, 0x64, 0x07,
	0x7f, 0x75, 0x3b, 0x86, 0x0e, 0x0d, 0xb2, 0xff, 0x8e, 0x9c, 0xfb, 0x76, 0x68, 0x2f, 0x94, 0xf6,
	0x74, 0x03, 0x2b, 0xf8, 0x27, 0x3b, 0x79, 0x88, 0x65, 0xaf, 0xf3, 0x0d, 0x76, 0x16, 0x0d, 0x9c,
	0xa1, 0xe7, 0x9c, 0xbe, 0xd1, 0x1b, 0xec, 0x6c, 0x6e, 0xe0, 0x1c, 0x3d, 0x27, 0x0e, 0x2e, 0x8c,
	0x37, 0xe2, 0x37, 0xe0, 0xf6, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x09, 0xf8, 0xf8, 0x03, 0x0e,
	0x08, 0x00, 0x00,
}
