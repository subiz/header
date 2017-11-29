// Code generated by protoc-gen-go. DO NOT EDIT.
// source: widget/widget.proto

/*
Package widget is a generated protocol buffer package.

It is generated from these files:
	widget/widget.proto

It has these top-level messages:
	Theme
	Sound
	Setting
	UserSetting
*/
package widget

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import account "bitbucket.org/subiz/header/account"
import common "bitbucket.org/subiz/header/common"
import user "bitbucket.org/subiz/header/user"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Event int32

const (
	Event_WidgetUserSettingReadRequested   Event = 0
	Event_WidgetUserSettingUpdateRequested Event = 1
	Event_WidgetSettingReadRequested       Event = 2
	Event_WidgetSettingUpdateRequested     Event = 3
)

var Event_name = map[int32]string{
	0: "WidgetUserSettingReadRequested",
	1: "WidgetUserSettingUpdateRequested",
	2: "WidgetSettingReadRequested",
	3: "WidgetSettingUpdateRequested",
}
var Event_value = map[string]int32{
	"WidgetUserSettingReadRequested":   0,
	"WidgetUserSettingUpdateRequested": 1,
	"WidgetSettingReadRequested":       2,
	"WidgetSettingUpdateRequested":     3,
}

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}
func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}
func (x *Event) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Event_value, data, "Event")
	if err != nil {
		return err
	}
	*x = Event(value)
	return nil
}
func (Event) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Theme_ButtonPosition int32

const (
	Theme_left  Theme_ButtonPosition = 0
	Theme_right Theme_ButtonPosition = 1
)

var Theme_ButtonPosition_name = map[int32]string{
	0: "left",
	1: "right",
}
var Theme_ButtonPosition_value = map[string]int32{
	"left":  0,
	"right": 1,
}

func (x Theme_ButtonPosition) Enum() *Theme_ButtonPosition {
	p := new(Theme_ButtonPosition)
	*p = x
	return p
}
func (x Theme_ButtonPosition) String() string {
	return proto.EnumName(Theme_ButtonPosition_name, int32(x))
}
func (x *Theme_ButtonPosition) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Theme_ButtonPosition_value, data, "Theme_ButtonPosition")
	if err != nil {
		return err
	}
	*x = Theme_ButtonPosition(value)
	return nil
}
func (Theme_ButtonPosition) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Theme_WindowMode int32

const (
	Theme_mini Theme_WindowMode = 0
	Theme_full Theme_WindowMode = 1
)

var Theme_WindowMode_name = map[int32]string{
	0: "mini",
	1: "full",
}
var Theme_WindowMode_value = map[string]int32{
	"mini": 0,
	"full": 1,
}

func (x Theme_WindowMode) Enum() *Theme_WindowMode {
	p := new(Theme_WindowMode)
	*p = x
	return p
}
func (x Theme_WindowMode) String() string {
	return proto.EnumName(Theme_WindowMode_name, int32(x))
}
func (x *Theme_WindowMode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Theme_WindowMode_value, data, "Theme_WindowMode")
	if err != nil {
		return err
	}
	*x = Theme_WindowMode(value)
	return nil
}
func (Theme_WindowMode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type Theme struct {
	BackgroundColor  *string `protobuf:"bytes,7,opt,name=background_color,json=backgroundColor" json:"background_color,omitempty"`
	CustomCss        *string `protobuf:"bytes,8,opt,name=custom_css,json=customCss" json:"custom_css,omitempty"`
	ButtonPosition   *string `protobuf:"bytes,3,opt,name=button_position,json=buttonPosition" json:"button_position,omitempty"`
	ButtonStyle      *string `protobuf:"bytes,4,opt,name=button_style,json=buttonStyle" json:"button_style,omitempty"`
	WindowMode       *string `protobuf:"bytes,5,opt,name=window_mode,json=windowMode" json:"window_mode,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Theme) Reset()                    { *m = Theme{} }
func (m *Theme) String() string            { return proto.CompactTextString(m) }
func (*Theme) ProtoMessage()               {}
func (*Theme) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Theme) GetBackgroundColor() string {
	if m != nil && m.BackgroundColor != nil {
		return *m.BackgroundColor
	}
	return ""
}

func (m *Theme) GetCustomCss() string {
	if m != nil && m.CustomCss != nil {
		return *m.CustomCss
	}
	return ""
}

func (m *Theme) GetButtonPosition() string {
	if m != nil && m.ButtonPosition != nil {
		return *m.ButtonPosition
	}
	return ""
}

func (m *Theme) GetButtonStyle() string {
	if m != nil && m.ButtonStyle != nil {
		return *m.ButtonStyle
	}
	return ""
}

func (m *Theme) GetWindowMode() string {
	if m != nil && m.WindowMode != nil {
		return *m.WindowMode
	}
	return ""
}

type Sound struct {
	Enabled           *bool   `protobuf:"varint,2,opt,name=enabled" json:"enabled,omitempty"`
	NewConversation   *string `protobuf:"bytes,3,opt,name=new_conversation,json=newConversation" json:"new_conversation,omitempty"`
	FileCreate        *string `protobuf:"bytes,4,opt,name=file_create,json=fileCreate" json:"file_create,omitempty"`
	NewMessage        *string `protobuf:"bytes,5,opt,name=new_message,json=newMessage" json:"new_message,omitempty"`
	MessageSendFailed *string `protobuf:"bytes,6,opt,name=message_send_failed,json=messageSendFailed" json:"message_send_failed,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *Sound) Reset()                    { *m = Sound{} }
func (m *Sound) String() string            { return proto.CompactTextString(m) }
func (*Sound) ProtoMessage()               {}
func (*Sound) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Sound) GetEnabled() bool {
	if m != nil && m.Enabled != nil {
		return *m.Enabled
	}
	return false
}

func (m *Sound) GetNewConversation() string {
	if m != nil && m.NewConversation != nil {
		return *m.NewConversation
	}
	return ""
}

func (m *Sound) GetFileCreate() string {
	if m != nil && m.FileCreate != nil {
		return *m.FileCreate
	}
	return ""
}

func (m *Sound) GetNewMessage() string {
	if m != nil && m.NewMessage != nil {
		return *m.NewMessage
	}
	return ""
}

func (m *Sound) GetMessageSendFailed() string {
	if m != nil && m.MessageSendFailed != nil {
		return *m.MessageSendFailed
	}
	return ""
}

type Setting struct {
	Ctx              *common.Context  `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	AccountId        *string          `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	Sound            *Sound           `protobuf:"bytes,4,opt,name=sound" json:"sound,omitempty"`
	Language         *string          `protobuf:"bytes,5,opt,name=language" json:"language,omitempty"`
	WelcomeMessage   *string          `protobuf:"bytes,6,opt,name=welcome_message,json=welcomeMessage" json:"welcome_message,omitempty"`
	Theme            *Theme           `protobuf:"bytes,7,opt,name=theme" json:"theme,omitempty"`
	Replytime        *int32           `protobuf:"varint,9,opt,name=replytime" json:"replytime,omitempty"`
	Agents           []*account.Agent `protobuf:"bytes,10,rep,name=agents" json:"agents,omitempty"`
	AgentIds         []string         `protobuf:"bytes,11,rep,name=agent_ids,json=agentIds" json:"agent_ids,omitempty"`
	CustomLanguage   *bool            `protobuf:"varint,12,opt,name=custom_language,json=customLanguage" json:"custom_language,omitempty"`
	LanguagePath     *string          `protobuf:"bytes,13,opt,name=language_path,json=languagePath" json:"language_path,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Setting) Reset()                    { *m = Setting{} }
func (m *Setting) String() string            { return proto.CompactTextString(m) }
func (*Setting) ProtoMessage()               {}
func (*Setting) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Setting) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *Setting) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *Setting) GetSound() *Sound {
	if m != nil {
		return m.Sound
	}
	return nil
}

func (m *Setting) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

func (m *Setting) GetWelcomeMessage() string {
	if m != nil && m.WelcomeMessage != nil {
		return *m.WelcomeMessage
	}
	return ""
}

func (m *Setting) GetTheme() *Theme {
	if m != nil {
		return m.Theme
	}
	return nil
}

func (m *Setting) GetReplytime() int32 {
	if m != nil && m.Replytime != nil {
		return *m.Replytime
	}
	return 0
}

func (m *Setting) GetAgents() []*account.Agent {
	if m != nil {
		return m.Agents
	}
	return nil
}

func (m *Setting) GetAgentIds() []string {
	if m != nil {
		return m.AgentIds
	}
	return nil
}

func (m *Setting) GetCustomLanguage() bool {
	if m != nil && m.CustomLanguage != nil {
		return *m.CustomLanguage
	}
	return false
}

func (m *Setting) GetLanguagePath() string {
	if m != nil && m.LanguagePath != nil {
		return *m.LanguagePath
	}
	return ""
}

type UserSetting struct {
	Ctx                 *common.Context  `protobuf:"bytes,1,opt,name=ctx" json:"ctx,omitempty"`
	Account             *account.Account `protobuf:"bytes,3,opt,name=account" json:"account,omitempty"`
	AccountId           *string          `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	User                *user.User       `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	UserId              *string          `protobuf:"bytes,7,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	SoundEnabled        *bool            `protobuf:"varint,5,opt,name=sound_enabled,json=soundEnabled" json:"sound_enabled,omitempty"`
	Language            *string          `protobuf:"bytes,6,opt,name=language" json:"language,omitempty"`
	SendTranscript      *bool            `protobuf:"varint,8,opt,name=send_transcript,json=sendTranscript" json:"send_transcript,omitempty"`
	AccountSetting      *Setting         `protobuf:"bytes,9,opt,name=account_setting,json=accountSetting" json:"account_setting,omitempty"`
	DesktopNotification *bool            `protobuf:"varint,10,opt,name=desktop_notification,json=desktopNotification" json:"desktop_notification,omitempty"`
	XXX_unrecognized    []byte           `json:"-"`
}

func (m *UserSetting) Reset()                    { *m = UserSetting{} }
func (m *UserSetting) String() string            { return proto.CompactTextString(m) }
func (*UserSetting) ProtoMessage()               {}
func (*UserSetting) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UserSetting) GetCtx() *common.Context {
	if m != nil {
		return m.Ctx
	}
	return nil
}

func (m *UserSetting) GetAccount() *account.Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *UserSetting) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *UserSetting) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UserSetting) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *UserSetting) GetSoundEnabled() bool {
	if m != nil && m.SoundEnabled != nil {
		return *m.SoundEnabled
	}
	return false
}

func (m *UserSetting) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

func (m *UserSetting) GetSendTranscript() bool {
	if m != nil && m.SendTranscript != nil {
		return *m.SendTranscript
	}
	return false
}

func (m *UserSetting) GetAccountSetting() *Setting {
	if m != nil {
		return m.AccountSetting
	}
	return nil
}

func (m *UserSetting) GetDesktopNotification() bool {
	if m != nil && m.DesktopNotification != nil {
		return *m.DesktopNotification
	}
	return false
}

func init() {
	proto.RegisterType((*Theme)(nil), "widget.Theme")
	proto.RegisterType((*Sound)(nil), "widget.Sound")
	proto.RegisterType((*Setting)(nil), "widget.Setting")
	proto.RegisterType((*UserSetting)(nil), "widget.UserSetting")
	proto.RegisterEnum("widget.Event", Event_name, Event_value)
	proto.RegisterEnum("widget.Theme_ButtonPosition", Theme_ButtonPosition_name, Theme_ButtonPosition_value)
	proto.RegisterEnum("widget.Theme_WindowMode", Theme_WindowMode_name, Theme_WindowMode_value)
}

func init() { proto.RegisterFile("widget/widget.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 776 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x8f, 0xdb, 0x44,
	0x14, 0x5d, 0x6f, 0xd6, 0x49, 0x7c, 0xbd, 0x9b, 0x84, 0x59, 0x24, 0xac, 0xa5, 0x2c, 0x6e, 0x96,
	0x8f, 0xb0, 0x0f, 0x0e, 0xe4, 0x89, 0x57, 0x88, 0x8a, 0x54, 0x89, 0xa2, 0xca, 0x69, 0xd5, 0x47,
	0xcb, 0xf1, 0xdc, 0x38, 0xa3, 0xb5, 0x67, 0x82, 0x67, 0xdc, 0xb4, 0xfc, 0x01, 0x1e, 0x78, 0xe4,
	0xff, 0xf0, 0xdb, 0xd0, 0x7c, 0x38, 0x49, 0x29, 0x2a, 0xe2, 0x25, 0xe3, 0x39, 0xf7, 0xcc, 0xcc,
	0x3d, 0xf7, 0x9e, 0x1b, 0xb8, 0xde, 0x33, 0x5a, 0xa2, 0x9a, 0xdb, 0x25, 0xd9, 0x35, 0x42, 0x09,
	0xd2, 0xb7, 0xbb, 0x9b, 0x6f, 0xd7, 0x4c, 0xad, 0xdb, 0xe2, 0x01, 0x55, 0x22, 0x9a, 0x72, 0x2e,
	0xdb, 0x35, 0xfb, 0x6d, 0xbe, 0xc5, 0x9c, 0x62, 0x33, 0xcf, 0x8b, 0x42, 0xb4, 0x5c, 0x75, 0xab,
	0x3d, 0x79, 0x93, 0x7c, 0xe0, 0x44, 0x21, 0xea, 0x5a, 0x70, 0xb7, 0x38, 0xfe, 0xfd, 0x07, 0xf8,
	0xad, 0x74, 0x3f, 0x96, 0x3b, 0xfd, 0xfd, 0x1c, 0xfc, 0x17, 0x5b, 0xac, 0x91, 0x7c, 0x03, 0x93,
	0x75, 0x5e, 0x3c, 0x94, 0x8d, 0x68, 0x39, 0xcd, 0x0a, 0x51, 0x89, 0x26, 0x1a, 0xc4, 0xde, 0x2c,
	0x48, 0xc7, 0x47, 0x7c, 0xa9, 0x61, 0xf2, 0x19, 0x40, 0xd1, 0x4a, 0x25, 0xea, 0xac, 0x90, 0x32,
	0x1a, 0x1a, 0x52, 0x60, 0x91, 0xa5, 0x94, 0xe4, 0x6b, 0x18, 0xaf, 0x5b, 0xa5, 0x04, 0xcf, 0x76,
	0x42, 0x32, 0xc5, 0x04, 0x8f, 0x7a, 0x86, 0x33, 0xb2, 0xf0, 0x73, 0x87, 0x92, 0xc7, 0x70, 0xe9,
	0x88, 0x52, 0xbd, 0xad, 0x30, 0xba, 0x30, 0xac, 0xd0, 0x62, 0x2b, 0x0d, 0x91, 0xcf, 0x21, 0xdc,
	0x33, 0x4e, 0xc5, 0x3e, 0xab, 0x05, 0xc5, 0xc8, 0x37, 0x0c, 0xb0, 0xd0, 0x33, 0x41, 0x71, 0xfa,
	0x25, 0x8c, 0x7e, 0x7c, 0xf7, 0xd6, 0x21, 0x5c, 0x54, 0xb8, 0x51, 0x93, 0x33, 0x12, 0x80, 0xdf,
	0xb0, 0x72, 0xab, 0x26, 0xde, 0x34, 0x06, 0x78, 0x75, 0x38, 0xa4, 0x29, 0x35, 0xe3, 0x6c, 0x72,
	0xa6, 0xbf, 0x36, 0x6d, 0x55, 0x4d, 0xbc, 0xe9, 0x5f, 0x1e, 0xf8, 0x2b, 0xad, 0x91, 0x44, 0x30,
	0x40, 0x9e, 0xaf, 0x2b, 0xa4, 0xd1, 0x79, 0xec, 0xcd, 0x86, 0x69, 0xb7, 0xd5, 0x35, 0xe2, 0xb8,
	0xcf, 0x0a, 0xc1, 0x5f, 0x63, 0x23, 0xf3, 0x13, 0x69, 0x63, 0x8e, 0xfb, 0xe5, 0x09, 0xac, 0x13,
	0xdf, 0xb0, 0x0a, 0xb3, 0xa2, 0xc1, 0x5c, 0x75, 0xd2, 0x40, 0x43, 0x4b, 0x83, 0x68, 0x82, 0xbe,
	0xab, 0x46, 0x29, 0xf3, 0xf2, 0xa0, 0x8c, 0xe3, 0xfe, 0x99, 0x45, 0x48, 0x02, 0xd7, 0x2e, 0x98,
	0x49, 0xe4, 0x34, 0xdb, 0xe4, 0x4c, 0xa7, 0xd4, 0x37, 0xc4, 0x8f, 0x5c, 0x68, 0x85, 0x9c, 0xfe,
	0x64, 0x02, 0xd3, 0x3f, 0x7a, 0x30, 0x58, 0xa1, 0x52, 0x8c, 0x97, 0xe4, 0x31, 0xf4, 0x0a, 0xf5,
	0x26, 0xf2, 0x62, 0x6f, 0x16, 0x2e, 0xc6, 0x89, 0xb3, 0xc7, 0x52, 0x70, 0x85, 0x6f, 0x54, 0xaa,
	0x63, 0xba, 0x89, 0xce, 0x66, 0x19, 0xb3, 0x42, 0x83, 0x34, 0x70, 0xc8, 0x53, 0x4a, 0xee, 0xc0,
	0x97, 0xba, 0x1a, 0x26, 0xf3, 0x70, 0x71, 0x95, 0x38, 0x33, 0x9b, 0x12, 0xa5, 0x36, 0x46, 0x6e,
	0x60, 0x58, 0xe5, 0xbc, 0x6c, 0x8f, 0x02, 0x0e, 0x7b, 0xed, 0x82, 0x3d, 0x56, 0x85, 0xa8, 0xf1,
	0xa0, 0xd1, 0xa6, 0x3e, 0x72, 0x70, 0xa7, 0xf3, 0x0e, 0x7c, 0xa5, 0x1d, 0x68, 0xdc, 0x76, 0xf2,
	0x92, 0xb1, 0x65, 0x6a, 0x63, 0xe4, 0x11, 0x04, 0x0d, 0xee, 0xaa, 0xb7, 0x8a, 0xd5, 0x18, 0x05,
	0xb1, 0x37, 0xf3, 0xd3, 0x23, 0x40, 0xbe, 0x82, 0x7e, 0x5e, 0x22, 0x57, 0x32, 0x82, 0xb8, 0x37,
	0x0b, 0x17, 0xa3, 0xa4, 0x9b, 0xa0, 0x1f, 0x34, 0x9c, 0xba, 0x28, 0xf9, 0x14, 0x02, 0xf3, 0x95,
	0x31, 0x2a, 0xa3, 0x30, 0xee, 0xe9, 0x84, 0x0d, 0xf0, 0x94, 0x1a, 0xdb, 0x3a, 0x57, 0x1f, 0x34,
	0x5d, 0x9a, 0xf6, 0x8f, 0x2c, 0xfc, 0x73, 0xa7, 0xec, 0x0e, 0xae, 0x3a, 0x46, 0xb6, 0xcb, 0xd5,
	0x36, 0xba, 0x32, 0xba, 0x2e, 0x3b, 0xf0, 0x79, 0xae, 0xb6, 0xba, 0x1b, 0xe1, 0x4b, 0x89, 0xcd,
	0xff, 0xe8, 0xc8, 0x3d, 0x0c, 0x5c, 0xda, 0xc6, 0x54, 0xe1, 0x62, 0x72, 0x94, 0x61, 0xd7, 0xb4,
	0x23, 0xfc, 0x57, 0xf7, 0x6e, 0xe1, 0x42, 0x0f, 0xb9, 0x6b, 0x1e, 0x24, 0x66, 0xe2, 0x75, 0x3a,
	0xa9, 0xc1, 0xc9, 0x27, 0x30, 0xd0, 0xab, 0x3e, 0x6b, 0x67, 0xbc, 0xaf, 0xb7, 0xa6, 0xed, 0x57,
	0xa6, 0xb5, 0x59, 0x37, 0x01, 0xbe, 0x29, 0xc1, 0xa5, 0x01, 0x9f, 0xb8, 0x31, 0x38, 0x6d, 0x7b,
	0xff, 0xfd, 0xb6, 0x1b, 0xb7, 0xaa, 0x26, 0xe7, 0xb2, 0x68, 0xd8, 0x4e, 0x99, 0x3f, 0x88, 0x61,
	0x3a, 0xd2, 0xf0, 0x8b, 0x03, 0x4a, 0xbe, 0x87, 0x71, 0xa7, 0x40, 0xda, 0x1a, 0x99, 0xbe, 0xea,
	0xe2, 0x74, 0x56, 0xb3, 0x70, 0x3a, 0x72, 0xbc, 0xae, 0x94, 0xdf, 0xc1, 0xc7, 0x14, 0xe5, 0x83,
	0x12, 0xbb, 0x8c, 0x0b, 0xc5, 0x36, 0xac, 0xb0, 0x93, 0x08, 0xe6, 0x9d, 0x6b, 0x17, 0xfb, 0xe5,
	0x24, 0x74, 0xff, 0xa7, 0x07, 0xfe, 0x93, 0xd7, 0xc8, 0x15, 0x99, 0xc2, 0xed, 0x2b, 0x73, 0xfd,
	0x49, 0x73, 0x52, 0xcc, 0x69, 0x8a, 0xbf, 0xb6, 0x28, 0x15, 0xd2, 0xc9, 0x19, 0xf9, 0x02, 0xe2,
	0xf7, 0x38, 0x2f, 0x77, 0x34, 0x57, 0x78, 0x64, 0x79, 0xe4, 0x16, 0x6e, 0x2c, 0xeb, 0x5f, 0x6f,
	0x39, 0x27, 0x31, 0x3c, 0x7a, 0x27, 0xfe, 0xcf, 0x1b, 0x7a, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x46, 0x87, 0x66, 0xbd, 0x28, 0x06, 0x00, 0x00,
}
