// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package client

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	auth "github.com/subiz/header/auth"
	common "github.com/subiz/header/common"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson19c08265DecodeGithubComSubizHeaderClient(in *jlexer.Lexer, out *Clients) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ctx":
			if in.IsNull() {
				in.Skip()
				out.Ctx = nil
			} else {
				if out.Ctx == nil {
					out.Ctx = new(common.Context)
				}
				easyjson19c08265DecodeGithubComSubizHeaderCommon(in, &*out.Ctx)
			}
		case "clients":
			if in.IsNull() {
				in.Skip()
				out.Clients = nil
			} else {
				in.Delim('[')
				if out.Clients == nil {
					if !in.IsDelim(']') {
						out.Clients = make([]*Client, 0, 8)
					} else {
						out.Clients = []*Client{}
					}
				} else {
					out.Clients = (out.Clients)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Client
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Client)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Clients = append(out.Clients, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson19c08265EncodeGithubComSubizHeaderClient(out *jwriter.Writer, in Clients) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Ctx != nil {
		const prefix string = ",\"ctx\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson19c08265EncodeGithubComSubizHeaderCommon(out, *in.Ctx)
	}
	if len(in.Clients) != 0 {
		const prefix string = ",\"clients\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Clients {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Clients) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson19c08265EncodeGithubComSubizHeaderClient(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Clients) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson19c08265EncodeGithubComSubizHeaderClient(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Clients) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson19c08265DecodeGithubComSubizHeaderClient(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Clients) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson19c08265DecodeGithubComSubizHeaderClient(l, v)
}
func easyjson19c08265DecodeGithubComSubizHeaderCommon(in *jlexer.Lexer, out *common.Context) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "event_id":
			out.EventId = string(in.String())
		case "state":
			if in.IsNull() {
				in.Skip()
				out.State = nil
			} else {
				out.State = in.Bytes()
			}
		case "node":
			out.Node = string(in.String())
		case "reply_topic":
			out.ReplyTopic = string(in.String())
		case "credential":
			if in.IsNull() {
				in.Skip()
				out.Credential = nil
			} else {
				if out.Credential == nil {
					out.Credential = new(auth.Credential)
				}
				(*out.Credential).UnmarshalEasyJSON(in)
			}
		case "tracing":
			if in.IsNull() {
				in.Skip()
				out.Tracing = nil
			} else {
				out.Tracing = in.Bytes()
			}
		case "reply_key":
			out.ReplyKey = string(in.String())
		case "by_device":
			if in.IsNull() {
				in.Skip()
				out.ByDevice = nil
			} else {
				if out.ByDevice == nil {
					out.ByDevice = new(common.Device)
				}
				easyjson19c08265DecodeGithubComSubizHeaderCommon1(in, &*out.ByDevice)
			}
		case "topic":
			out.Topic = string(in.String())
		case "partition":
			out.Partition = int32(in.Int32())
		case "offset":
			out.Offset = int64(in.Int64())
		case "term":
			out.Term = uint64(in.Uint64())
		case "router_topic":
			out.RouterTopic = string(in.String())
		case "idempotency_key":
			out.IdempotencyKey = string(in.String())
		case "env":
			out.Env = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson19c08265EncodeGithubComSubizHeaderCommon(out *jwriter.Writer, in common.Context) {
	out.RawByte('{')
	first := true
	_ = first
	if in.EventId != "" {
		const prefix string = ",\"event_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.EventId))
	}
	if len(in.State) != 0 {
		const prefix string = ",\"state\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.State)
	}
	if in.Node != "" {
		const prefix string = ",\"node\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Node))
	}
	if in.ReplyTopic != "" {
		const prefix string = ",\"reply_topic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ReplyTopic))
	}
	if in.Credential != nil {
		const prefix string = ",\"credential\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Credential).MarshalEasyJSON(out)
	}
	if len(in.Tracing) != 0 {
		const prefix string = ",\"tracing\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Base64Bytes(in.Tracing)
	}
	if in.ReplyKey != "" {
		const prefix string = ",\"reply_key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ReplyKey))
	}
	if in.ByDevice != nil {
		const prefix string = ",\"by_device\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson19c08265EncodeGithubComSubizHeaderCommon1(out, *in.ByDevice)
	}
	if in.Topic != "" {
		const prefix string = ",\"topic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Topic))
	}
	if in.Partition != 0 {
		const prefix string = ",\"partition\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Partition))
	}
	if in.Offset != 0 {
		const prefix string = ",\"offset\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Offset))
	}
	if in.Term != 0 {
		const prefix string = ",\"term\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Term))
	}
	if in.RouterTopic != "" {
		const prefix string = ",\"router_topic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RouterTopic))
	}
	if in.IdempotencyKey != "" {
		const prefix string = ",\"idempotency_key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.IdempotencyKey))
	}
	if in.Env != "" {
		const prefix string = ",\"env\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Env))
	}
	out.RawByte('}')
}
func easyjson19c08265DecodeGithubComSubizHeaderCommon1(in *jlexer.Lexer, out *common.Device) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ip":
			out.Ip = string(in.String())
		case "user_agent":
			out.UserAgent = string(in.String())
		case "screen_resolution":
			out.ScreenResolution = string(in.String())
		case "timezone":
			out.Timezone = string(in.String())
		case "language":
			out.Language = string(in.String())
		case "referrer":
			out.Referrer = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "platform":
			out.Platform = string(in.String())
		case "source_referrer":
			out.SourceReferrer = string(in.String())
		case "ga_tracking_ids":
			if in.IsNull() {
				in.Skip()
				out.GaTrackingIds = nil
			} else {
				in.Delim('[')
				if out.GaTrackingIds == nil {
					if !in.IsDelim(']') {
						out.GaTrackingIds = make([]string, 0, 4)
					} else {
						out.GaTrackingIds = []string{}
					}
				} else {
					out.GaTrackingIds = (out.GaTrackingIds)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.GaTrackingIds = append(out.GaTrackingIds, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson19c08265EncodeGithubComSubizHeaderCommon1(out *jwriter.Writer, in common.Device) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Ip != "" {
		const prefix string = ",\"ip\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Ip))
	}
	if in.UserAgent != "" {
		const prefix string = ",\"user_agent\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UserAgent))
	}
	if in.ScreenResolution != "" {
		const prefix string = ",\"screen_resolution\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ScreenResolution))
	}
	if in.Timezone != "" {
		const prefix string = ",\"timezone\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Timezone))
	}
	if in.Language != "" {
		const prefix string = ",\"language\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Language))
	}
	if in.Referrer != "" {
		const prefix string = ",\"referrer\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Referrer))
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	if in.Platform != "" {
		const prefix string = ",\"platform\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Platform))
	}
	if in.SourceReferrer != "" {
		const prefix string = ",\"source_referrer\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SourceReferrer))
	}
	if len(in.GaTrackingIds) != 0 {
		const prefix string = ",\"ga_tracking_ids\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v11, v12 := range in.GaTrackingIds {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson19c08265DecodeGithubComSubizHeaderClient1(in *jlexer.Lexer, out *Client) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ctx":
			if in.IsNull() {
				in.Skip()
				out.Ctx = nil
			} else {
				if out.Ctx == nil {
					out.Ctx = new(common.Context)
				}
				easyjson19c08265DecodeGithubComSubizHeaderCommon(in, &*out.Ctx)
			}
		case "id":
			if in.IsNull() {
				in.Skip()
				out.Id = nil
			} else {
				if out.Id == nil {
					out.Id = new(string)
				}
				*out.Id = string(in.String())
			}
		case "secret":
			if in.IsNull() {
				in.Skip()
				out.Secret = nil
			} else {
				if out.Secret == nil {
					out.Secret = new(string)
				}
				*out.Secret = string(in.String())
			}
		case "logo_url":
			if in.IsNull() {
				in.Skip()
				out.LogoUrl = nil
			} else {
				if out.LogoUrl == nil {
					out.LogoUrl = new(string)
				}
				*out.LogoUrl = string(in.String())
			}
		case "account_id":
			if in.IsNull() {
				in.Skip()
				out.AccountId = nil
			} else {
				if out.AccountId == nil {
					out.AccountId = new(string)
				}
				*out.AccountId = string(in.String())
			}
		case "is_verified":
			if in.IsNull() {
				in.Skip()
				out.IsVerified = nil
			} else {
				if out.IsVerified == nil {
					out.IsVerified = new(bool)
				}
				*out.IsVerified = bool(in.Bool())
			}
		case "verified":
			if in.IsNull() {
				in.Skip()
				out.Verified = nil
			} else {
				if out.Verified == nil {
					out.Verified = new(int64)
				}
				*out.Verified = int64(in.Int64())
			}
		case "redirect_uri":
			if in.IsNull() {
				in.Skip()
				out.RedirectUri = nil
			} else {
				if out.RedirectUri == nil {
					out.RedirectUri = new(string)
				}
				*out.RedirectUri = string(in.String())
			}
		case "type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				if out.Type == nil {
					out.Type = new(string)
				}
				*out.Type = string(in.String())
			}
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		case "version":
			if in.IsNull() {
				in.Skip()
				out.Version = nil
			} else {
				if out.Version == nil {
					out.Version = new(string)
				}
				*out.Version = string(in.String())
			}
		case "is_enabled":
			if in.IsNull() {
				in.Skip()
				out.IsEnabled = nil
			} else {
				if out.IsEnabled == nil {
					out.IsEnabled = new(bool)
				}
				*out.IsEnabled = bool(in.Bool())
			}
		case "created":
			if in.IsNull() {
				in.Skip()
				out.Created = nil
			} else {
				if out.Created == nil {
					out.Created = new(int64)
				}
				*out.Created = int64(in.Int64())
			}
		case "modified":
			if in.IsNull() {
				in.Skip()
				out.Modified = nil
			} else {
				if out.Modified == nil {
					out.Modified = new(int64)
				}
				*out.Modified = int64(in.Int64())
			}
		case "webhook_uri":
			if in.IsNull() {
				in.Skip()
				out.WebhookUri = nil
			} else {
				if out.WebhookUri == nil {
					out.WebhookUri = new(string)
				}
				*out.WebhookUri = string(in.String())
			}
		case "events":
			if in.IsNull() {
				in.Skip()
				out.Events = nil
			} else {
				in.Delim('[')
				if out.Events == nil {
					if !in.IsDelim(']') {
						out.Events = make([]string, 0, 4)
					} else {
						out.Events = []string{}
					}
				} else {
					out.Events = (out.Events)[:0]
				}
				for !in.IsDelim(']') {
					var v13 string
					v13 = string(in.String())
					out.Events = append(out.Events, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "channel_type":
			if in.IsNull() {
				in.Skip()
				out.ChannelType = nil
			} else {
				if out.ChannelType == nil {
					out.ChannelType = new(string)
				}
				*out.ChannelType = string(in.String())
			}
		case "availability_uri":
			if in.IsNull() {
				in.Skip()
				out.AvailabilityUri = nil
			} else {
				if out.AvailabilityUri == nil {
					out.AvailabilityUri = new(string)
				}
				*out.AvailabilityUri = string(in.String())
			}
		case "ping_uri":
			if in.IsNull() {
				in.Skip()
				out.PingUri = nil
			} else {
				if out.PingUri == nil {
					out.PingUri = new(string)
				}
				*out.PingUri = string(in.String())
			}
		case "is_internal":
			if in.IsNull() {
				in.Skip()
				out.IsInternal = nil
			} else {
				if out.IsInternal == nil {
					out.IsInternal = new(bool)
				}
				*out.IsInternal = bool(in.Bool())
			}
		case "unsubscribe_uri":
			if in.IsNull() {
				in.Skip()
				out.UnsubscribeUri = nil
			} else {
				if out.UnsubscribeUri == nil {
					out.UnsubscribeUri = new(string)
				}
				*out.UnsubscribeUri = string(in.String())
			}
		case "scopes":
			if in.IsNull() {
				in.Skip()
				out.Scopes = nil
			} else {
				in.Delim('[')
				if out.Scopes == nil {
					if !in.IsDelim(']') {
						out.Scopes = make([]string, 0, 4)
					} else {
						out.Scopes = []string{}
					}
				} else {
					out.Scopes = (out.Scopes)[:0]
				}
				for !in.IsDelim(']') {
					var v14 string
					v14 = string(in.String())
					out.Scopes = append(out.Scopes, v14)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "bot_default_job_title":
			if in.IsNull() {
				in.Skip()
				out.BotDefaultJobTitle = nil
			} else {
				if out.BotDefaultJobTitle == nil {
					out.BotDefaultJobTitle = new(string)
				}
				*out.BotDefaultJobTitle = string(in.String())
			}
		case "bot_default_fullname":
			if in.IsNull() {
				in.Skip()
				out.BotDefaultFullname = nil
			} else {
				if out.BotDefaultFullname == nil {
					out.BotDefaultFullname = new(string)
				}
				*out.BotDefaultFullname = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson19c08265EncodeGithubComSubizHeaderClient1(out *jwriter.Writer, in Client) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Ctx != nil {
		const prefix string = ",\"ctx\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson19c08265EncodeGithubComSubizHeaderCommon(out, *in.Ctx)
	}
	if in.Id != nil {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Id))
	}
	if in.Secret != nil {
		const prefix string = ",\"secret\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Secret))
	}
	if in.LogoUrl != nil {
		const prefix string = ",\"logo_url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.LogoUrl))
	}
	if in.AccountId != nil {
		const prefix string = ",\"account_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.AccountId))
	}
	if in.IsVerified != nil {
		const prefix string = ",\"is_verified\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(*in.IsVerified))
	}
	if in.Verified != nil {
		const prefix string = ",\"verified\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(*in.Verified))
	}
	if in.RedirectUri != nil {
		const prefix string = ",\"redirect_uri\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.RedirectUri))
	}
	if in.Type != nil {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Type))
	}
	if in.Name != nil {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Name))
	}
	if in.Version != nil {
		const prefix string = ",\"version\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Version))
	}
	if in.IsEnabled != nil {
		const prefix string = ",\"is_enabled\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(*in.IsEnabled))
	}
	if in.Created != nil {
		const prefix string = ",\"created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(*in.Created))
	}
	if in.Modified != nil {
		const prefix string = ",\"modified\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(*in.Modified))
	}
	if in.WebhookUri != nil {
		const prefix string = ",\"webhook_uri\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.WebhookUri))
	}
	if len(in.Events) != 0 {
		const prefix string = ",\"events\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v15, v16 := range in.Events {
				if v15 > 0 {
					out.RawByte(',')
				}
				out.String(string(v16))
			}
			out.RawByte(']')
		}
	}
	if in.ChannelType != nil {
		const prefix string = ",\"channel_type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.ChannelType))
	}
	if in.AvailabilityUri != nil {
		const prefix string = ",\"availability_uri\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.AvailabilityUri))
	}
	if in.PingUri != nil {
		const prefix string = ",\"ping_uri\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.PingUri))
	}
	if in.IsInternal != nil {
		const prefix string = ",\"is_internal\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(*in.IsInternal))
	}
	if in.UnsubscribeUri != nil {
		const prefix string = ",\"unsubscribe_uri\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.UnsubscribeUri))
	}
	if len(in.Scopes) != 0 {
		const prefix string = ",\"scopes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v17, v18 := range in.Scopes {
				if v17 > 0 {
					out.RawByte(',')
				}
				out.String(string(v18))
			}
			out.RawByte(']')
		}
	}
	if in.BotDefaultJobTitle != nil {
		const prefix string = ",\"bot_default_job_title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.BotDefaultJobTitle))
	}
	if in.BotDefaultFullname != nil {
		const prefix string = ",\"bot_default_fullname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.BotDefaultFullname))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Client) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson19c08265EncodeGithubComSubizHeaderClient1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Client) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson19c08265EncodeGithubComSubizHeaderClient1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Client) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson19c08265DecodeGithubComSubizHeaderClient1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Client) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson19c08265DecodeGithubComSubizHeaderClient1(l, v)
}
func easyjson19c08265DecodeGithubComSubizHeaderClient2(in *jlexer.Lexer, out *AuthorizedClient) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "ctx":
			if in.IsNull() {
				in.Skip()
				out.Ctx = nil
			} else {
				if out.Ctx == nil {
					out.Ctx = new(common.Context)
				}
				easyjson19c08265DecodeGithubComSubizHeaderCommon(in, &*out.Ctx)
			}
		case "client":
			if in.IsNull() {
				in.Skip()
				out.Client = nil
			} else {
				if out.Client == nil {
					out.Client = new(Client)
				}
				(*out.Client).UnmarshalEasyJSON(in)
			}
		case "issue_account_id":
			if in.IsNull() {
				in.Skip()
				out.IssueAccountId = nil
			} else {
				if out.IssueAccountId == nil {
					out.IssueAccountId = new(string)
				}
				*out.IssueAccountId = string(in.String())
			}
		case "issuer":
			if in.IsNull() {
				in.Skip()
				out.Issuer = nil
			} else {
				if out.Issuer == nil {
					out.Issuer = new(string)
				}
				*out.Issuer = string(in.String())
			}
		case "scopes":
			if in.IsNull() {
				in.Skip()
				out.Scopes = nil
			} else {
				in.Delim('[')
				if out.Scopes == nil {
					if !in.IsDelim(']') {
						out.Scopes = make([]string, 0, 4)
					} else {
						out.Scopes = []string{}
					}
				} else {
					out.Scopes = (out.Scopes)[:0]
				}
				for !in.IsDelim(']') {
					var v19 string
					v19 = string(in.String())
					out.Scopes = append(out.Scopes, v19)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "kind":
			if in.IsNull() {
				in.Skip()
				out.Kind = nil
			} else {
				if out.Kind == nil {
					out.Kind = new(string)
				}
				*out.Kind = string(in.String())
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson19c08265EncodeGithubComSubizHeaderClient2(out *jwriter.Writer, in AuthorizedClient) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Ctx != nil {
		const prefix string = ",\"ctx\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson19c08265EncodeGithubComSubizHeaderCommon(out, *in.Ctx)
	}
	if in.Client != nil {
		const prefix string = ",\"client\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Client).MarshalEasyJSON(out)
	}
	if in.IssueAccountId != nil {
		const prefix string = ",\"issue_account_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.IssueAccountId))
	}
	if in.Issuer != nil {
		const prefix string = ",\"issuer\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Issuer))
	}
	if len(in.Scopes) != 0 {
		const prefix string = ",\"scopes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v20, v21 := range in.Scopes {
				if v20 > 0 {
					out.RawByte(',')
				}
				out.String(string(v21))
			}
			out.RawByte(']')
		}
	}
	if in.Kind != nil {
		const prefix string = ",\"kind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Kind))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AuthorizedClient) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson19c08265EncodeGithubComSubizHeaderClient2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AuthorizedClient) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson19c08265EncodeGithubComSubizHeaderClient2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AuthorizedClient) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson19c08265DecodeGithubComSubizHeaderClient2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AuthorizedClient) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson19c08265DecodeGithubComSubizHeaderClient2(l, v)
}
