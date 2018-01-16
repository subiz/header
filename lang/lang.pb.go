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
	T_invalid_left                      T = 14
	T_internal_error                    T = 30
	T_invalid_input                     T = 22
	T_invalid_form                      T = 20
	T_access_token_expired              T = 21
	T_invalid_credential                T = 13
	T_credential_not_set                T = 7
	T_wrong_account_in_credential       T = 8
	T_wrong_password                    T = 18
	T_wrong_user_in_credential          T = 10
	T_invalid_agent_state               T = 19
	T_access_deny                       T = 9
	T_unable_to_send_message            T = 11
	T_topic_is_empty                    T = 12
	T_invalid_json                      T = 15
	T_invalid_protobuf                  T = 16
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
)

var T_name = map[int32]string{
	0:   "undefined",
	1:   "user_has_already_in_conversation",
	2:   "conversation_closed",
	3:   "invalid_invite",
	4:   "invalid_agent",
	5:   "member_is_not_in_conversation",
	6:   "conversation_not_found",
	14:  "invalid_left",
	30:  "internal_error",
	22:  "invalid_input",
	20:  "invalid_form",
	21:  "access_token_expired",
	13:  "invalid_credential",
	7:   "credential_not_set",
	8:   "wrong_account_in_credential",
	18:  "wrong_password",
	10:  "wrong_user_in_credential",
	19:  "invalid_agent_state",
	9:   "access_deny",
	11:  "unable_to_send_message",
	12:  "topic_is_empty",
	15:  "invalid_json",
	16:  "invalid_protobuf",
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
}
var T_value = map[string]int32{
	"undefined":                         0,
	"user_has_already_in_conversation":  1,
	"conversation_closed":               2,
	"invalid_invite":                    3,
	"invalid_agent":                     4,
	"member_is_not_in_conversation":     5,
	"conversation_not_found":            6,
	"invalid_left":                      14,
	"internal_error":                    30,
	"invalid_input":                     22,
	"invalid_form":                      20,
	"access_token_expired":              21,
	"invalid_credential":                13,
	"credential_not_set":                7,
	"wrong_account_in_credential":       8,
	"wrong_password":                    18,
	"wrong_user_in_credential":          10,
	"invalid_agent_state":               19,
	"access_deny":                       9,
	"unable_to_send_message":            11,
	"topic_is_empty":                    12,
	"invalid_json":                      15,
	"invalid_protobuf":                  16,
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
	// 1003 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x55, 0x59, 0x8f, 0x1c, 0x35,
	0x10, 0x4e, 0x6f, 0xb2, 0x9b, 0xac, 0xf7, 0xaa, 0xf5, 0x1e, 0x6c, 0x36, 0x90, 0x10, 0x01, 0x12,
	0x2c, 0x90, 0x00, 0x11, 0x6f, 0xdc, 0x0f, 0xf0, 0xc2, 0x03, 0x47, 0x00, 0x89, 0xcb, 0xf2, 0xb4,
	0x6b, 0x66, 0x9c, 0x71, 0xdb, 0x2d, 0xdb, 0x3d, 0xb3, 0xc3, 0x2f, 0x44, 0xe2, 0x92, 0xe0, 0x95,
	0x1b, 0xc4, 0xf1, 0x1b, 0x78, 0x41, 0xe5, 0xee, 0xde, 0xf6, 0x06, 0x5e, 0xa6, 0xa7, 0x3f, 0xd7,
	0xf1, 0xd5, 0x57, 0x55, 0x6e, 0xb6, 0x63, 0xa4, 0x9d, 0xdc, 0xa6, 0x9f, 0x5b, 0xb5, 0x77, 0xd1,
	0xf1, 0x4b, 0xf4, 0xff, 0xe4, 0x9f, 0x6d, 0x56, 0xdc, 0xe5, 0x5b, 0x6c, 0xbd, 0xb1, 0x0a, 0xc7,
	0xda, 0xa2, 0x82, 0x0b, 0xfc, 0x51, 0xf6, 0x70, 0x13, 0xd0, 0x8b, 0xa9, 0x0c, 0x42, 0x1a, 0x8f,
	0x52, 0x2d, 0x85, 0xb6, 0xa2, 0x74, 0x76, 0x8e, 0x3e, 0xc8, 0xa8, 0x9d, 0x85, 0x82, 0x3f, 0xc0,
	0xf6, 0x72, 0x44, 0x94, 0xc6, 0x05, 0x54, 0xb0, 0xc2, 0x39, 0xdb, 0xd6, 0x76, 0x2e, 0x8d, 0x56,
	0x42, 0xdb, 0xb9, 0x8e, 0x08, 0x17, 0xf9, 0x2e, 0xdb, 0xea, 0x31, 0x39, 0x41, 0x1b, 0xe1, 0x12,
	0xbf, 0xc9, 0x1e, 0xaa, 0xb0, 0x1a, 0xa1, 0x17, 0x3a, 0x08, 0xeb, 0xe2, 0x7f, 0x52, 0xac, 0xf2,
	0x63, 0x76, 0x78, 0x2e, 0x05, 0x59, 0x8d, 0x5d, 0x63, 0x15, 0xac, 0x71, 0x60, 0x9b, 0x7d, 0x44,
	0x83, 0xe3, 0x08, 0xdb, 0x6d, 0xde, 0x88, 0xde, 0x4a, 0x23, 0xd0, 0x7b, 0xe7, 0xe1, 0x7a, 0x9e,
	0x57, 0xdb, 0xba, 0x89, 0x70, 0x98, 0x3b, 0x8e, 0x9d, 0xaf, 0x60, 0x9f, 0x1f, 0xb1, 0x7d, 0x59,
	0x96, 0x18, 0x82, 0x88, 0x6e, 0x86, 0x56, 0xe0, 0x69, 0xad, 0x3d, 0x2a, 0x38, 0xe0, 0x87, 0x8c,
	0xf7, 0xb6, 0xa5, 0x47, 0x85, 0x36, 0x6a, 0x69, 0x60, 0x8b, 0xf0, 0xe1, 0x3d, 0xd1, 0x0a, 0x18,
	0xe1, 0x32, 0xbf, 0xc1, 0xae, 0x2d, 0xbc, 0xb3, 0x13, 0x21, 0xcb, 0xd2, 0x35, 0xb6, 0xad, 0x69,
	0x70, 0xbc, 0x42, 0x1c, 0x5b, 0x83, 0x5a, 0x86, 0xb0, 0x70, 0x5e, 0x01, 0xe7, 0x0f, 0xb2, 0xa3,
	0x16, 0x4b, 0xa2, 0x9f, 0xf7, 0x60, 0x24, 0xf3, 0x39, 0xe5, 0x44, 0x88, 0x32, 0x22, 0xec, 0xf1,
	0x1d, 0xb6, 0xd1, 0xb1, 0x56, 0x68, 0x97, 0xb0, 0x4e, 0x6a, 0x35, 0x56, 0x8e, 0x0c, 0x8a, 0xe8,
	0x44, 0x40, 0xab, 0x44, 0x85, 0x21, 0xc8, 0x09, 0xc2, 0x06, 0xe5, 0x8d, 0xae, 0xd6, 0x25, 0x69,
	0x8d, 0x55, 0x1d, 0x97, 0xb0, 0x99, 0x0b, 0x71, 0x2f, 0x38, 0x0b, 0x3b, 0x7c, 0x9f, 0x41, 0x8f,
	0xa4, 0x21, 0x19, 0x35, 0x63, 0x00, 0xf2, 0x1d, 0xe2, 0x1a, 0x57, 0xce, 0xe0, 0x71, 0xbe, 0xce,
	0x56, 0xdb, 0x30, 0x4f, 0xf0, 0x6d, 0xc6, 0x5a, 0xfa, 0x71, 0x59, 0x23, 0x9c, 0xe4, 0x84, 0x67,
	0x72, 0x3c, 0x93, 0x22, 0x25, 0x86, 0x27, 0x29, 0x8e, 0x92, 0x51, 0x8e, 0x64, 0xc0, 0xae, 0x3f,
	0x4f, 0xf1, 0x0d, 0x76, 0x39, 0xea, 0x0a, 0x5d, 0x13, 0xe1, 0x69, 0xbe, 0xc7, 0x76, 0x16, 0x38,
	0x0a, 0xae, 0x9c, 0x61, 0xec, 0x2c, 0x6e, 0x51, 0x99, 0x6d, 0x98, 0x16, 0xb8, 0x9d, 0xb7, 0x34,
	0xb5, 0x0b, 0x9e, 0xe1, 0x07, 0x6c, 0xb7, 0x17, 0x7c, 0x18, 0x91, 0x67, 0x29, 0x5e, 0x2b, 0xd9,
	0x00, 0x3e, 0x97, 0xbb, 0x63, 0x25, 0xb5, 0x81, 0x17, 0x88, 0x58, 0x6d, 0x64, 0x3e, 0x5e, 0x2f,
	0xf2, 0xab, 0xec, 0xa0, 0xf5, 0x9d, 0x78, 0xd7, 0xd4, 0xd9, 0xd1, 0x4b, 0x94, 0x2d, 0xd5, 0x2e,
	0x4a, 0xa3, 0xc9, 0x22, 0xd5, 0xfd, 0x32, 0x2d, 0x51, 0x0b, 0x37, 0xde, 0xc0, 0x2b, 0x24, 0x4b,
	0xfb, 0x6a, 0x65, 0x85, 0xf0, 0x2a, 0x69, 0xdb, 0xd9, 0x0f, 0xb1, 0x5e, 0x23, 0x36, 0xad, 0x55,
	0xc7, 0x1f, 0x5e, 0xe7, 0xd7, 0xd9, 0xf1, 0xd9, 0xcc, 0xe5, 0xc3, 0xdf, 0xf6, 0xfd, 0x8d, 0x7c,
	0x26, 0xbb, 0xfe, 0x0a, 0xad, 0xe0, 0xad, 0xbc, 0x9d, 0x95, 0x0c, 0x33, 0x78, 0x9b, 0xea, 0xf7,
	0xd2, 0x2a, 0x57, 0xe9, 0xcf, 0x7a, 0xc5, 0xdf, 0xe1, 0x8f, 0xb1, 0x9b, 0xaa, 0xa9, 0x8d, 0x2e,
	0x65, 0xc4, 0x21, 0x82, 0xc7, 0x12, 0xf5, 0x1c, 0x55, 0x67, 0xf6, 0x2e, 0x11, 0xb3, 0x18, 0x17,
	0xce, 0xcf, 0x3a, 0xe8, 0x2e, 0xd5, 0xed, 0x83, 0x14, 0x33, 0x5c, 0x66, 0x25, 0xbc, 0x47, 0xea,
	0xdd, 0x5b, 0x44, 0x11, 0xf4, 0xc4, 0x76, 0xa6, 0xef, 0x53, 0xb1, 0x68, 0xe7, 0xc4, 0x7f, 0xac,
	0x27, 0x1d, 0xfa, 0x01, 0x31, 0x0c, 0xa5, 0x5f, 0xd6, 0x7d, 0x73, 0x3f, 0x4c, 0x77, 0xc8, 0x54,
	0x1a, 0x83, 0x76, 0x82, 0x59, 0xd0, 0x8f, 0xc8, 0xb4, 0x5b, 0x24, 0x1b, 0x16, 0xe8, 0xe1, 0x63,
	0x32, 0x0d, 0xe8, 0xe7, 0xe8, 0x85, 0xd1, 0x21, 0x62, 0x9f, 0xeb, 0x13, 0xea, 0x54, 0x17, 0x75,
	0xac, 0x4d, 0x1e, 0xe5, 0xd3, 0xf3, 0xa3, 0x42, 0x43, 0x28, 0x88, 0xed, 0x7d, 0x66, 0x65, 0x1a,
	0x70, 0x5a, 0xbd, 0x01, 0x53, 0x43, 0x37, 0x2b, 0xf5, 0x3c, 0x20, 0x89, 0x9e, 0xe6, 0xb6, 0xed,
	0x48, 0x5f, 0xc0, 0x98, 0x6f, 0xb2, 0x2b, 0xe1, 0x4e, 0xf7, 0x36, 0x49, 0x17, 0xc9, 0x22, 0x64,
	0xfb, 0xdb, 0x9d, 0x4c, 0x29, 0x05, 0x9d, 0xa4, 0xad, 0x6c, 0x31, 0xcd, 0x8f, 0xd8, 0x1e, 0x1a,
	0x19, 0xa2, 0x2e, 0x03, 0x4a, 0x5f, 0x4e, 0xbb, 0x83, 0xcf, 0x0b, 0x7e, 0x30, 0xec, 0x61, 0xc4,
	0xaa, 0x36, 0xd4, 0xf8, 0x2f, 0x0a, 0x72, 0xa0, 0x00, 0x13, 0xaf, 0x55, 0x1e, 0xe9, 0xcb, 0x82,
	0xdf, 0x60, 0xc7, 0x8b, 0xa9, 0x8e, 0x48, 0xd2, 0x08, 0xe5, 0x2a, 0xa9, 0xf3, 0x69, 0xfe, 0xaa,
	0xe0, 0xd7, 0xd8, 0xe1, 0xc8, 0xc8, 0x72, 0x96, 0x0c, 0x74, 0x3e, 0xcf, 0x5f, 0x77, 0x87, 0xb4,
	0x75, 0x4a, 0xdc, 0xa7, 0xc3, 0x37, 0x05, 0xbf, 0xca, 0xf6, 0xb3, 0x71, 0x8c, 0x67, 0x03, 0xff,
	0x6d, 0xa2, 0x59, 0x4b, 0x1f, 0xb0, 0x15, 0xbe, 0x25, 0xf3, 0x5d, 0x0a, 0x37, 0xdc, 0xb9, 0x11,
	0x27, 0xbe, 0x9d, 0x5f, 0xad, 0xe0, 0xfb, 0x54, 0xc3, 0xff, 0x1c, 0xc2, 0x0f, 0x05, 0x3f, 0x64,
	0xbb, 0x0b, 0x1c, 0x4d, 0x9d, 0x9b, 0x65, 0x04, 0x7e, 0x2c, 0xf8, 0x1e, 0xdb, 0x26, 0x11, 0xb2,
	0x1c, 0x3f, 0xa5, 0xd4, 0x0b, 0xaf, 0xe3, 0xb9, 0xd4, 0x3f, 0x27, 0x38, 0x7d, 0x86, 0x72, 0xf8,
	0x97, 0x94, 0x14, 0x4f, 0xb1, 0x6c, 0x22, 0x8a, 0x30, 0x45, 0xd3, 0xb7, 0xe5, 0xd7, 0x82, 0x03,
	0xdb, 0x38, 0xab, 0x2e, 0x04, 0xf8, 0xad, 0xe0, 0xbb, 0xc3, 0x1a, 0x4d, 0x2b, 0x59, 0xc2, 0xef,
	0x09, 0x2a, 0x9d, 0x0d, 0x4d, 0xef, 0xf7, 0x47, 0xc1, 0x39, 0xdb, 0xaa, 0xe4, 0x69, 0xa5, 0xad,
	0x52, 0x23, 0x42, 0xe1, 0xcf, 0x82, 0x1f, 0xb3, 0x83, 0x4c, 0x29, 0xa5, 0x53, 0xd5, 0x33, 0x5c,
	0xc2, 0x5f, 0x05, 0xdf, 0x67, 0x3b, 0x67, 0x37, 0xab, 0x4b, 0xec, 0xe0, 0xef, 0xe2, 0xe4, 0x11,
	0x56, 0xbc, 0xc9, 0xd7, 0xd8, 0x0a, 0x5a, 0xb8, 0x40, 0xcf, 0x39, 0x7d, 0x57, 0xd7, 0xd8, 0x4a,
	0x69, 0xe1, 0x22, 0x3d, 0xc7, 0x1e, 0x56, 0x47, 0x6b, 0xe9, 0x2a, 0xbe, 0xf3, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x26, 0xf7, 0xf3, 0xe6, 0xc2, 0x07, 0x00, 0x00,
}
