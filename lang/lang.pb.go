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
	// 801 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x54, 0x4b, 0x93, 0xdc, 0x34,
	0x10, 0xce, 0x6c, 0xb2, 0x9b, 0xac, 0xf6, 0xd5, 0xd1, 0x3e, 0x58, 0x02, 0x04, 0x52, 0x40, 0x15,
	0x2c, 0x90, 0x00, 0x29, 0x6e, 0xbc, 0x0f, 0x70, 0xe1, 0xc0, 0x23, 0x40, 0x15, 0x2f, 0x95, 0x46,
	0x6a, 0x7b, 0x15, 0xcb, 0x92, 0x4b, 0x92, 0x67, 0x18, 0x7e, 0x1f, 0x07, 0xaa, 0xe0, 0x47, 0x51,
	0x6d, 0xd9, 0x58, 0x43, 0x2e, 0xe3, 0x51, 0xab, 0xfb, 0xeb, 0xaf, 0xbf, 0xee, 0x16, 0x3b, 0xb1,
	0xd2, 0xd5, 0x8f, 0xe8, 0xe7, 0x61, 0x17, 0x7c, 0xf2, 0xfc, 0x16, 0xfd, 0xbf, 0xfa, 0xf3, 0x80,
	0x2d, 0x9e, 0xf0, 0x23, 0xb6, 0xdf, 0x3b, 0x8d, 0x95, 0x71, 0xa8, 0xe1, 0x06, 0x7f, 0x8d, 0xbd,
	0xd2, 0x47, 0x0c, 0xe2, 0x5a, 0x46, 0x21, 0x6d, 0x40, 0xa9, 0x37, 0xc2, 0x38, 0xa1, 0xbc, 0x5b,
	0x61, 0x88, 0x32, 0x19, 0xef, 0x60, 0xc1, 0x9f, 0x63, 0xa7, 0xa5, 0x45, 0x28, 0xeb, 0x23, 0x6a,
	0xd8, 0xe1, 0x9c, 0x1d, 0x1b, 0xb7, 0x92, 0xd6, 0x68, 0x61, 0xdc, 0xca, 0x24, 0x84, 0x9b, 0xfc,
	0x2e, 0x3b, 0x9a, 0x6c, 0xb2, 0x46, 0x97, 0xe0, 0x16, 0x7f, 0xc0, 0x5e, 0x6a, 0xb1, 0x5d, 0x62,
	0x10, 0x26, 0x0a, 0xe7, 0xd3, 0x33, 0x29, 0x76, 0xf9, 0x3d, 0x76, 0xb1, 0x95, 0x82, 0xbc, 0x2a,
	0xdf, 0x3b, 0x0d, 0x7b, 0x1c, 0xd8, 0xe1, 0x84, 0x68, 0xb1, 0x4a, 0x70, 0x9c, 0xf3, 0x26, 0x0c,
	0x4e, 0x5a, 0x81, 0x21, 0xf8, 0x00, 0xf7, 0xcb, 0xbc, 0xc6, 0x75, 0x7d, 0x82, 0x8b, 0x32, 0xb0,
	0xf2, 0xa1, 0x85, 0x33, 0x7e, 0xc9, 0xce, 0xa4, 0x52, 0x18, 0xa3, 0x48, 0xbe, 0x41, 0x27, 0xf0,
	0xf7, 0xce, 0x04, 0xd4, 0x70, 0xce, 0x2f, 0x18, 0x9f, 0x7c, 0x55, 0x40, 0x8d, 0x2e, 0x19, 0x69,
	0xe1, 0x88, 0xec, 0xf3, 0x79, 0xa0, 0x15, 0x31, 0xc1, 0x6d, 0xfe, 0x32, 0x7b, 0x61, 0x1d, 0xbc,
	0xab, 0x85, 0x54, 0xca, 0xf7, 0x2e, 0xd7, 0x34, 0x07, 0xde, 0x21, 0x8e, 0xd9, 0xa1, 0x93, 0x31,
	0xae, 0x7d, 0xd0, 0xc0, 0xf9, 0x8b, 0xec, 0x32, 0xdb, 0x06, 0xd1, 0xb7, 0x23, 0x18, 0xc9, 0xbc,
	0xa5, 0x9c, 0x88, 0x49, 0x26, 0x84, 0x53, 0x7e, 0xc2, 0x0e, 0x46, 0xd6, 0x1a, 0xdd, 0x06, 0xf6,
	0x49, 0xad, 0xde, 0xc9, 0xa5, 0x45, 0x91, 0xbc, 0x88, 0xe8, 0xb4, 0x68, 0x31, 0x46, 0x59, 0x23,
	0x1c, 0x50, 0xde, 0xe4, 0x3b, 0xa3, 0x48, 0x6b, 0x6c, 0xbb, 0xb4, 0x81, 0xc3, 0x52, 0x88, 0xa7,
	0xd1, 0x3b, 0x38, 0xe1, 0x67, 0x0c, 0x26, 0xcb, 0x30, 0x24, 0xcb, 0xbe, 0x02, 0xa0, 0xd8, 0x19,
	0xd7, 0x7a, 0xd5, 0xc0, 0x1b, 0x7c, 0x9f, 0xed, 0x66, 0x98, 0x37, 0xf9, 0x31, 0x63, 0x99, 0x7e,
	0xda, 0x74, 0x08, 0x57, 0x25, 0xe1, 0x46, 0x56, 0x8d, 0x14, 0x43, 0x62, 0x78, 0x8b, 0x70, 0xb4,
	0x4c, 0x72, 0x29, 0x23, 0x8e, 0xfd, 0x79, 0x9b, 0x1f, 0xb0, 0xdb, 0xc9, 0xb4, 0xe8, 0xfb, 0x04,
	0xef, 0xf0, 0x53, 0x76, 0xb2, 0xc6, 0x65, 0xf4, 0xaa, 0xc1, 0x34, 0x7a, 0x3c, 0xa4, 0x32, 0x33,
	0x4c, 0x36, 0x3c, 0x2a, 0x5b, 0x3a, 0xb4, 0x0b, 0xde, 0xe5, 0xe7, 0xec, 0xee, 0x24, 0xf8, 0x3c,
	0x22, 0xef, 0x11, 0x5e, 0x96, 0x6c, 0x36, 0xbe, 0x5f, 0x86, 0x63, 0x2b, 0x8d, 0x85, 0x0f, 0x89,
	0x58, 0x67, 0x65, 0x39, 0x5e, 0x1f, 0xf1, 0xe7, 0xd9, 0x79, 0x8e, 0xad, 0x83, 0xef, 0xbb, 0xe2,
	0xea, 0x63, 0xca, 0x36, 0xd4, 0x2e, 0x94, 0x35, 0xe4, 0x31, 0xd4, 0xfd, 0x09, 0x2d, 0x51, 0x36,
	0xf7, 0xc1, 0xc2, 0xa7, 0x24, 0x4b, 0x3e, 0x3a, 0xd9, 0x22, 0x7c, 0x46, 0xda, 0x8e, 0xfe, 0x33,
	0xd6, 0xe7, 0xc4, 0x26, 0x7b, 0x8d, 0xfc, 0xe1, 0x0b, 0x7e, 0x9f, 0xdd, 0xfb, 0x6f, 0xe6, 0xca,
	0xe1, 0xcf, 0x7d, 0xff, 0xb2, 0x9c, 0xc9, 0xb1, 0xbf, 0xc2, 0x68, 0xf8, 0xba, 0x6c, 0x67, 0x2b,
	0x63, 0x03, 0xdf, 0x50, 0xfd, 0x41, 0x3a, 0xed, 0x5b, 0xf3, 0xc7, 0xa4, 0xf8, 0xb7, 0xfc, 0x75,
	0xf6, 0x40, 0xf7, 0x9d, 0x35, 0x4a, 0x26, 0x9c, 0x11, 0x02, 0x2a, 0x34, 0x2b, 0xd4, 0xa3, 0xdb,
	0x77, 0x44, 0xcc, 0x61, 0x5a, 0xfb, 0xd0, 0x8c, 0xa6, 0x27, 0x54, 0x77, 0x88, 0x52, 0x34, 0xb8,
	0x29, 0x4a, 0xf8, 0x9e, 0xd4, 0x7b, 0xba, 0x4e, 0x22, 0x9a, 0xda, 0x8d, 0xae, 0x3f, 0x50, 0xb1,
	0xe8, 0x56, 0xc4, 0xbf, 0x32, 0xf5, 0x68, 0xfd, 0x91, 0x18, 0x46, 0x15, 0x36, 0xdd, 0xd4, 0xdc,
	0x9f, 0x86, 0x37, 0xe4, 0x5a, 0x5a, 0x8b, 0xae, 0xc6, 0x02, 0xf4, 0x67, 0x72, 0x1d, 0x17, 0xc9,
	0xc5, 0x35, 0x06, 0xf8, 0x85, 0x5c, 0x23, 0x86, 0x15, 0x06, 0x61, 0x4d, 0x4c, 0x38, 0xe5, 0xfa,
	0x95, 0x3a, 0x35, 0xa2, 0x56, 0xc6, 0x96, 0x28, 0xbf, 0x6d, 0x8f, 0x0a, 0x0d, 0xa1, 0x20, 0xb6,
	0xff, 0x73, 0x53, 0xc3, 0x80, 0xd3, 0xea, 0xcd, 0x36, 0x3d, 0x77, 0xb3, 0xd5, 0x1f, 0x00, 0x92,
	0xe8, 0xc3, 0xdc, 0xe6, 0x8e, 0x4c, 0x05, 0x54, 0xfc, 0x90, 0xdd, 0x89, 0x8f, 0xc7, 0x53, 0x3d,
	0x3c, 0x24, 0xeb, 0x58, 0xec, 0xef, 0x78, 0x73, 0x4d, 0x29, 0xe8, 0x66, 0xd8, 0xca, 0x6c, 0x33,
	0xfc, 0x92, 0x9d, 0xa2, 0x95, 0x31, 0x19, 0x15, 0x51, 0x06, 0x75, 0x3d, 0x5e, 0xfc, 0xb5, 0xe0,
	0xe7, 0xf3, 0x1e, 0x26, 0x6c, 0x3b, 0x4b, 0x8d, 0xff, 0x7b, 0x41, 0x01, 0x04, 0x50, 0x07, 0xa3,
	0x4b, 0xa4, 0x7f, 0x16, 0x57, 0xaf, 0xb2, 0xc5, 0x57, 0x7c, 0x8f, 0xed, 0xa0, 0x83, 0x1b, 0xf4,
	0x5d, 0xd1, 0x03, 0xbd, 0xc7, 0x76, 0x94, 0x83, 0x9b, 0xf4, 0xad, 0x02, 0xec, 0x2e, 0xf7, 0x86,
	0x9d, 0x7e, 0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x6e, 0x15, 0x83, 0x0b, 0x06, 0x00,
	0x00,
}
