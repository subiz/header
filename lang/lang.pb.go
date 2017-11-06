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
	T_wrong_user_in_credential          T = 10
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
)

var T_name = map[int32]string{
	0:  "undefined",
	1:  "user_has_already_in_conversation",
	2:  "conversation_closed",
	3:  "invalid_invite",
	4:  "invalid_agent",
	5:  "member_is_not_in_conversation",
	6:  "conversation_not_found",
	14: "invalid_left",
	30: "internal_error",
	22: "invalid_input",
	20: "invalid_form",
	21: "access_token_expired",
	13: "invalid_credential",
	7:  "credential_not_set",
	8:  "wrong_account_in_credential",
	10: "wrong_user_in_credential",
	9:  "access_deny",
	11: "unable_to_send_message",
	12: "topic_is_empty",
	15: "invalid_json",
	16: "invalid_protobuf",
	40: "unable_to_lock",
	41: "empty",
	42: "wrong_type",
	43: "invalid_kafka_topic",
	44: "database_error",
	45: "timeout",
	46: "websocket_error",
	47: "kafka_error",
	48: "invalid_token",
	49: "account_not_found",
	50: "agent_not_found",
	60: "invalid_email",
	61: "plan_not_found",
	62: "agent_group_not_found",
	63: "empty_client_type",
	64: "empty_url",
	65: "empty_name",
	66: "client_not_found",
	70: "empty_account",
	71: "invalid_conversation_state",
	80: "invalid_message_id",
	81: "invalid_mask",
	82: "randomize_error",
	83: "duplicated_message_received_error",
	84: "network_error",
	85: "rsa_key_not_found",
	86: "jwt_sign_error",
	87: "env_config_error",
	90: "scrypt_error",
	91: "challenge_not_found",
	92: "wrong_answer",
	93: "server_listen_error",
	94: "scrypt_file_not_found",
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
	"wrong_user_in_credential":          10,
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
}

func (x T) String() string {
	return proto.EnumName(T_name, int32(x))
}
func (T) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("lang.T", T_name, T_value)
}

func init() { proto.RegisterFile("lang/lang.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 662 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x54, 0x69, 0x6f, 0xdc, 0x36,
	0x10, 0xad, 0x5b, 0x1f, 0x5d, 0xfa, 0xd8, 0x31, 0x7d, 0xd4, 0x75, 0x5b, 0xb7, 0x06, 0x5a, 0xa0,
	0x71, 0x12, 0x3b, 0xc7, 0xd7, 0xdc, 0x1f, 0x92, 0xaf, 0x39, 0x9c, 0x04, 0xc8, 0x45, 0x70, 0xc5,
	0x91, 0x4c, 0x8b, 0x22, 0x05, 0x92, 0xda, 0xcd, 0xe6, 0x47, 0xe5, 0x37, 0x06, 0xa3, 0x23, 0x4b,
	0x23, 0x5f, 0x0c, 0x6b, 0xf4, 0xe6, 0xcd, 0xbc, 0xf7, 0x46, 0xcb, 0xc6, 0x46, 0xda, 0xe2, 0x8c,
	0xfe, 0x9c, 0xd6, 0xde, 0x45, 0xc7, 0x97, 0xe9, 0xff, 0x93, 0xaf, 0x23, 0xb6, 0x74, 0xce, 0x37,
	0xd9, 0xa8, 0xb1, 0x0a, 0x73, 0x6d, 0x51, 0xc1, 0x4f, 0xfc, 0x5f, 0xf6, 0x4f, 0x13, 0xd0, 0x8b,
	0x0b, 0x19, 0x84, 0x34, 0x1e, 0xa5, 0x9a, 0x0b, 0x6d, 0x45, 0xe6, 0xec, 0x14, 0x7d, 0x90, 0x51,
	0x3b, 0x0b, 0x4b, 0xfc, 0x37, 0xb6, 0x93, 0x56, 0x44, 0x66, 0x5c, 0x40, 0x05, 0x3f, 0x73, 0xce,
	0xb6, 0xb4, 0x9d, 0x4a, 0xa3, 0x95, 0xd0, 0x76, 0xaa, 0x23, 0xc2, 0x2f, 0x7c, 0x9b, 0x6d, 0x0e,
	0x35, 0x59, 0xa0, 0x8d, 0xb0, 0xcc, 0x8f, 0xd9, 0x5f, 0x15, 0x56, 0x13, 0xf4, 0x42, 0x07, 0x61,
	0x5d, 0xfc, 0x61, 0xc4, 0x0a, 0x3f, 0x64, 0xfb, 0x57, 0x46, 0x10, 0x2a, 0x77, 0x8d, 0x55, 0xb0,
	0xca, 0x81, 0x6d, 0x0c, 0x8c, 0x06, 0xf3, 0x08, 0x5b, 0xdd, 0xdc, 0x88, 0xde, 0x4a, 0x23, 0xd0,
	0x7b, 0xe7, 0xe1, 0x28, 0x9d, 0xab, 0x6d, 0xdd, 0x44, 0xd8, 0x4f, 0x1b, 0x73, 0xe7, 0x2b, 0xd8,
	0xe5, 0x07, 0x6c, 0x57, 0x66, 0x19, 0x86, 0x20, 0xa2, 0x2b, 0xd1, 0x0a, 0xfc, 0x5c, 0x6b, 0x8f,
	0x0a, 0xf6, 0xf8, 0x3e, 0xe3, 0x03, 0x36, 0xf3, 0xa8, 0xd0, 0x46, 0x2d, 0x0d, 0x6c, 0x52, 0x7d,
	0xf1, 0xdc, 0xae, 0x15, 0x30, 0xc2, 0x1a, 0xff, 0x9b, 0xfd, 0x31, 0xf3, 0xce, 0x16, 0x42, 0x66,
	0x99, 0x6b, 0x6c, 0xa7, 0x69, 0xd1, 0xf8, 0x2b, 0xff, 0x93, 0x1d, 0x74, 0x80, 0xd6, 0xe0, 0xab,
	0x6f, 0x19, 0x1f, 0xb3, 0xf5, 0x7e, 0x11, 0x85, 0x76, 0x0e, 0x23, 0x32, 0xa0, 0xb1, 0x72, 0x62,
	0x50, 0x44, 0x27, 0x02, 0x5a, 0x25, 0x2a, 0x0c, 0x41, 0x16, 0x08, 0xeb, 0x24, 0x37, 0xba, 0x5a,
	0x67, 0x64, 0x1f, 0x56, 0x75, 0x9c, 0xc3, 0x46, 0xaa, 0xed, 0x32, 0x38, 0x0b, 0x63, 0xbe, 0xcb,
	0x60, 0xa8, 0xb4, 0xb9, 0x4f, 0x9a, 0x1c, 0x80, 0x7a, 0x17, 0xbc, 0xc6, 0x65, 0x25, 0xfc, 0xcf,
	0x47, 0x6c, 0xa5, 0xa3, 0xb9, 0xc6, 0xb7, 0x18, 0xeb, 0xb6, 0x8c, 0xf3, 0x1a, 0xe1, 0x84, 0xa2,
	0x1e, 0x48, 0x4a, 0x99, 0x97, 0x52, 0xb4, 0x83, 0xe1, 0x3a, 0xf1, 0x28, 0x19, 0xe5, 0x44, 0x06,
	0xec, 0x2d, 0xbf, 0xc1, 0xd7, 0xd9, 0x5a, 0xd4, 0x15, 0xba, 0x26, 0xc2, 0x4d, 0xbe, 0xc3, 0xc6,
	0x33, 0x9c, 0x04, 0x97, 0x95, 0x18, 0x7b, 0xc4, 0x29, 0xc9, 0xec, 0x68, 0xba, 0xc2, 0x59, 0x9a,
	0x52, 0x9b, 0x00, 0xdc, 0xe2, 0x7b, 0x6c, 0x7b, 0xf0, 0x70, 0x91, 0xfa, 0x6d, 0xe2, 0x6b, 0xef,
	0x27, 0x29, 0xde, 0x49, 0xdb, 0xb1, 0x92, 0xda, 0xc0, 0x3d, 0x5a, 0xac, 0x36, 0x32, 0xbd, 0x98,
	0xfb, 0xfc, 0x77, 0xb6, 0xd7, 0xf5, 0x16, 0xde, 0x35, 0x75, 0xf2, 0xea, 0x01, 0x4d, 0x6b, 0xb5,
	0x8b, 0xcc, 0x68, 0x42, 0xb4, 0xba, 0x1f, 0xd2, 0x77, 0xd1, 0x95, 0x1b, 0x6f, 0xe0, 0x11, 0xd9,
	0xd2, 0x3d, 0x5a, 0x59, 0x21, 0x3c, 0x26, 0x6f, 0x7b, 0xfc, 0x82, 0xeb, 0x09, 0x6d, 0xd3, 0xa1,
	0xfa, 0xfd, 0xe1, 0x29, 0x3f, 0x62, 0x87, 0xdf, 0xcf, 0x28, 0xbd, 0xe7, 0x10, 0x65, 0x44, 0x78,
	0x96, 0x9e, 0x59, 0x9f, 0xaf, 0xd0, 0x0a, 0x9e, 0xa7, 0x71, 0x56, 0x32, 0x94, 0xf0, 0x82, 0xf4,
	0x7b, 0x69, 0x95, 0xab, 0xf4, 0x97, 0xc1, 0xf1, 0x97, 0xfc, 0x3f, 0x76, 0xac, 0x9a, 0xda, 0xe8,
	0x4c, 0x46, 0x5c, 0x30, 0x78, 0xcc, 0x50, 0x4f, 0x51, 0xf5, 0xb0, 0x57, 0xb4, 0x98, 0xc5, 0x38,
	0x73, 0xbe, 0xec, 0x4b, 0xe7, 0xa4, 0xdb, 0x07, 0x29, 0x4a, 0x9c, 0x27, 0x12, 0x5e, 0x93, 0x7b,
	0x97, 0xb3, 0x28, 0x82, 0x2e, 0x6c, 0x0f, 0x7d, 0x43, 0x62, 0xd1, 0x4e, 0x69, 0xff, 0x5c, 0x17,
	0x7d, 0xf5, 0x2d, 0x6d, 0x18, 0x32, 0x3f, 0xaf, 0x87, 0x70, 0xdf, 0xb5, 0x3f, 0x0b, 0x17, 0xd2,
	0x18, 0xb4, 0x05, 0x26, 0xa4, 0xef, 0x09, 0xda, 0x7f, 0x1b, 0x36, 0xcc, 0xd0, 0xc3, 0x07, 0x82,
	0x06, 0xf4, 0x53, 0xf4, 0xc2, 0xe8, 0x10, 0x71, 0x98, 0xf5, 0x91, 0x92, 0xea, 0x59, 0x73, 0x6d,
	0x52, 0x96, 0x4f, 0x93, 0xd5, 0xf6, 0x8a, 0xef, 0x7e, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x12,
	0xc9, 0xaa, 0xd0, 0x04, 0x00, 0x00,
}
