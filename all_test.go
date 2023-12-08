package header

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"math"
	"sort"
	"strings"
	"testing"
	"time"

	pb "github.com/subiz/header/account"
	cpb "github.com/subiz/header/common"
	"google.golang.org/protobuf/proto"
)

func TestObjectPath(t *testing.T) {
	testcases := []struct {
		in   any
		path string
		val  any
	}{
		//{nil, "", nil},
		//{nil, "abc", nil},
		//{map[string]any{"user": map[string]string{"fullname": "thanh"}}, "user.fullname", "thanh"},
		{map[string]any{"user": []any{map[string]string{"fullname": "thanh"}}}, "user.0.fullname", "thanh"},
		//{map[string]any{"user": []any{nil, map[string]string{"fullname": "thanh"}}}, "user.1.fullname", "thanh"},
	}

	for i, tc := range testcases {
		interf := ObjectPath(tc.in, tc.path)
		if interf == nil && tc.val == nil {
			continue
		}

		if interf != nil && tc.val == nil || interf == nil && tc.val != nil {
			t.Errorf("WRONG AT TEST #%d, expect %v, got %v", i+1, tc.val, interf)
			continue
		}

		// !nil
		if fmt.Sprintf("%v", interf) != fmt.Sprintf("%v", tc.val) {
			t.Errorf("WRONG AT TEST #%d, expect %s, got %s", i+1, tc.val, interf)
		}
	}
}

func TestNormEmail(t *testing.T) {
	testcases := []struct {
		in  string
		out string
	}{
		{"", ""},
		{"thanhpk@live.com, thanhpk@live.com 't@.com'", "thanhpk@live.com,thanhpk@live.com,t@.com"},
	}

	for i, tc := range testcases {
		if tc.out != NormEmail(tc.in) {
			t.Errorf("WRONG AT TEST #%d, expect %s, got %s", i+1, tc.out, NormPhone(tc.in))
		}
	}
}

func TestEmailAddress(t *testing.T) {
	testcases := []struct {
		in  string
		out string
	}{
		{"", ""},
		{"thanhpk@live.com, thanhpk@live.com 't@.com'", "thanhpk@live.com"},
		{"thanhpk@live.com,xxx@gmail.com", "thanhpk@live.com"},
	}

	for i, tc := range testcases {
		if tc.out != EmailAddress(tc.in) {
			t.Errorf("WRONG AT TEST #%d, expect %s, got %s", i+1, tc.out, NormPhone(tc.in))
		}
	}
}

func TestNormPhone(t *testing.T) {
	testcases := []struct {
		in  string
		out string
	}{
		{"", ""},
		{"(84)35 9423 423", "0359423423"},
		{"(+84)35.9423 423", "0359423423"},
		{"(84)35 9423 423,+843630985", "0359423423,03630985"},
		{",,,", ""},
		{",,123123123,", "123123123"},
		{",,,", ""},
		{"abc@gmail.com", ""},
		{"245abc@gmail.com", ""},
		{",;84123123123,", "0123123123"},
	}

	for i, tc := range testcases {
		if tc.out != NormPhone(tc.in) {
			t.Errorf("WRONG AT TEST #%d, expect %s, got %s", i+1, tc.out, NormPhone(tc.in))
		}
	}
}

func TestPhoneNumber(t *testing.T) {
	testcases := []struct {
		in  string
		out string
	}{
		{"", ""},
		{"(84)35 9423 423", "84359423423"},
		{"(+84)35.9423 423", "84359423423"},
		{"(84)35 9423 423,+843630985", "84359423423"},
		{",,,", ""},
		{",,123123123,", "123123123"},
		{",,,", ""},
		{"abc@gmail.com,45", ""},
		{"245abc@gmail.com", "245"},
		{",;84123123123,", "84123123123"},
	}

	for i, tc := range testcases {
		if tc.out != PhoneNumber(tc.in) {
			t.Errorf("WRONG AT TEST #%d, expect %s, got %s", i+1, tc.out, PhoneNumber(tc.in))
		}
	}
}

func TestAssignObject(t *testing.T) {
	dst := &PhoneDevice{Id: "1", Name: "dst", Created: 1}
	src := &PhoneDevice{Id: "2", Name: "src", Created: 2}
	AssignObject(dst, src, []string{"name"})

	if dst.Id != "1" || dst.Name != "src" || dst.Created != 1 {
		t.Error("should be eq")
	}

	AssignObject(dst, nil, []string{"name"})
	src = nil
	AssignObject(dst, src, []string{"name"})

	if dst.Id != "1" || dst.Name != "src" || dst.Created != 1 {
		t.Error("should be eq 2")
	}
}

func TestAssignAgent(t *testing.T) {
	pS := func(s string) *string {
		return &s
	}

	dst := &pb.Agent{Id: pS("1"), Fullname: pS("dst"), AvatarUrl: pS("http://dst.com")}
	src := &pb.Agent{Id: pS("1"), Fullname: pS("src"), AvatarUrl: pS("http://src.com")}
	AssignObject(dst, src, []string{"avatar_url"})
	fmt.Println("D", dst)
}

func TestAssignByte(t *testing.T) {
	dst := &PhoneDevice{Id: "1", Name: "dst", Created: 1}
	AssignJSONByte(dst, []byte(`{"id": "4", "created": 5}`))

	if dst.Id != "4" || dst.Name != "dst" || dst.Created != 5 {
		t.Error("should be eq")
	}
}

func TestPartition(t *testing.T) {
	fmt.Println("PAR", Fnv32("usrtgmzbsbqcunkvjrhuq")%50)
	shardNumber := int(crc32.ChecksumIEEE([]byte("acriviayfmabzskstrpq"))) % 4
	fmt.Println(shardNumber)
}

func TestUnpack(t *testing.T) {
	str := "0x1215757372766b696c6b71706c737877696e736a76776c40c196c289c0314a0e636f6e74656e745f766965776564a20120321e320838343737363037323a0838343838393534324a0838343737363037329203471a15757372766b696c6b71706c737877696e736a76776c4a2e18d3fb84a30730c8e7c42838e28ea32848a2d0af065087dcb028580160e4bdbc286887dcb0287088a9b628800102"
	str = strings.TrimPrefix(str, "0x")

	bs, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}

	ev := &Event{}
	proto.Unmarshal(bs, ev)
	out, _ := json.Marshal(ev)

	fmt.Println(string(out))
	// b, _ := proto.Marshal(ev)
	// str = hex.EncodeToString(b)

	// fmt.Println("STR", string(str))
}

func TestLang(t *testing.T) {
	str := &I18NString{
		En_US: "tieng anh",
		Vi_VN: "tieng viet",
		Ja_JP: "tieng nhat",
	}

	if GetI18n(str, "en-US", "ja-JP") != "tieng anh" {
		t.Fatalf("should be tieng anh, got %s", GetI18n(str, "en-US", "jp-JP"))
	}

	if GetI18n(str, "ar-MA", "ja-JP") != "tieng nhat" {
		t.Fatalf("should be tieng nhat, got %s.", GetI18n(str, "ar-MA", "jp-JP"))
	}
}

func TestAllLang(t *testing.T) {
	str := &I18NString{
		En_US: "tieng anh",
		Vi_VN: "tieng viet",
		Ja_JP: "tieng nhat",
	}

	a := GetAllI18ns(str)
	fmt.Println(a)
}

func TestPerm(t *testing.T) {
	fmt.Println("D", checkAccess([]string{"all"}, "phone_device:w"))
	fmt.Println("A", checkAccess([]string{"view_others", "agent", "account_setting"}, "phone_device:w"))
}

func TestUserViewCondition(t *testing.T) {
	// submit ít nhất 2 form dang-ky trong vòng 5 ngày (trường company-size trong form phải > 20 người)
	// không chạy với cus id là 243
	// với những khách trong khoảng 23/4/2022 12:00 UTC+7 tới 21/4/2023 12:00 UTC+7
	// không xem trang google.com (tieu de Google) ít nhất 2 lần trong 2 ngày gần đây
	// co conversation tag la xyz, ghy
	condition := &UserViewCondition{
		All: []*UserViewCondition{{
			Id:  "1",
			Key: "event:form_submitted",
			Event: &EventCondition{
				UiType:  "form_submitted",
				Inverse: false,
				AtLeast: 2,
				Created: &DatetimeCondition{
					Last: 432000, // 5 days
				},
			},
			All: []*UserViewCondition{{
				Key:  "data.form.name",
				Text: &TextCondition{Eq: []string{"dang-ky"}},
			}, {
				Key:    "data.form.fields#(key=\"company-size\"]).first",
				Number: &FloatCondition{Gt: 20},
			}},
		}, {
			Id:   "2",
			Key:  "attr:cus_id",
			Text: &TextCondition{Neq: []string{"243"}},
		}, {
			Id:       "3",
			Key:      "attr:created",
			Datetime: &DatetimeCondition{Between: []int64{429423, 43434}},
		}, {
			Id:  "4",
			Key: "event:content_viewed",
			Event: &EventCondition{
				UiType:  "page",
				Inverse: true, // khong xem trang
				AtLeast: 2,
				Created: &DatetimeCondition{
					Last: 172800, // 2 days in secs
				},
			},
			All: []*UserViewCondition{{
				Key:  "data.product.title",
				Text: &TextCondition{Eq: []string{"Google"}},
			}, {
				Key:  "data.product.url",
				Text: &TextCondition{Contain: []string{"google.com"}, CaseSensitive: false},
			}},
		}, {
			Id:  "5",
			Key: "conversation_tag",
			Text: &TextCondition{
				Contain: []string{"234234", "3222234"},
			},
		}, {
			Id:  "6",
			Key: "order",
			Text: &TextCondition{
				Contain: []string{"234234", "3222234"},
			},
		}, {
			Id:   "7",
			Key:  "segment",
			Text: &TextCondition{Neq: []string{"243"}},
		}},
	}

	b, _ := json.Marshal(condition)
	fmt.Println("OUT", string(b))
}

var DEFS = map[string]*AttributeDefinition{
	"fullname": &AttributeDefinition{Name: "Fullname", Key: "fullname", Type: "text", IsSystem: true},
	"banned":   &AttributeDefinition{Name: "Banned", Key: "is_ban", Type: "boolean", IsSystem: true, IsReadonly: true},
	"emails":   &AttributeDefinition{Name: "Email", Key: "emails", Type: "text", IsSystem: true},
	"phones":   &AttributeDefinition{Name: "Phone", Key: "phones", Type: "text", IsSystem: true},
	"lifecycle_stage": {
		Type:   "text",
		Select: "dropdown",
		Items: []*AttributeDefinitionListItem{
			{Value: "subscriber", I18NLabel: &I18NString{En_US: "Subscriber", Vi_VN: "Đã theo dõi"}, Label: "Đã theo dõi"},
			{Value: "lead", I18NLabel: &I18NString{En_US: "Lead", Vi_VN: "Tiềm năng"}, Label: "Tiềm năng"},
			{Value: "trial", I18NLabel: &I18NString{En_US: "Trial", Vi_VN: "Dùng thử"}, Label: "Dùng thử"},
			{Value: "qualified lead", I18NLabel: &I18NString{En_US: "Sale Qualified", Vi_VN: "Chất lượng"}, Label: "Chất lượng"},
			{Value: "opportunity", I18NLabel: &I18NString{En_US: "Opportunity", Vi_VN: "Cơ hội"}, Label: "Cơ hội"},
			{Value: "customer", I18NLabel: &I18NString{En_US: "Customer", Vi_VN: "Khách hàng"}, Label: "Khách hàng"},
		},
		Key: "lifecycle_stage", IsSystem: true,
	},
	"created":   {Name: "Created", Type: "datetime", Key: "created", IsSystem: true, IsReadonly: true},
	"rank":      {Name: "Rank", Type: "number", Key: "rank", IsSystem: true},
	"record_id": {Name: "Record ID", Type: "text", Key: "record_id", IsSystem: true},
	"address":   {Name: "Address", Type: "text", Key: "address", PreventAutoOverride: true},
}

func isEqualAttr(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)

	ajs, _ := json.Marshal(a)
	bjs, _ := json.Marshal(b)
	return string(ajs) == string(bjs)
}

func isEqualUserAttribute(usera, userb []*Attribute) string {
	if len(usera) != len(userb) {
		return "attribute len"
	}
	for _, a := range usera {
		found := false
		for _, b := range userb {
			if a.Key != b.Key {
				continue
			}
			found = true
			if a.Modified != b.Modified {
				return "modified"
			}

			if a.By != b.By {
				return "by"
			}

			if a.ByType != b.ByType {
				return "bytype"
			}

			if a.Text != b.Text {
				return "text"
			}

			if a.Datetime != b.Datetime {
				return "datetime"
			}

			if a.Boolean != b.Boolean {
				return "boolean"
			}

			if math.Abs(a.Number-b.Number) > 0.00001 {
				return "number"
			}

			if a.UserValue != b.UserValue {
				return "uservalue"
			}
			if a.ConnectorValue != b.ConnectorValue {
				return "connectorvalue"
			}

			if a.Action != b.Action {
				return "action"
			}

			if !isEqualAttr(a.OtherLabels, b.OtherLabels) {
				return "other labels"
			}

			if !isEqualAttr(a.OtherValues, b.OtherValues) {
				return "other values"
			}
		}

		if !found {
			return "not found"
		}
	}
	return ""
}

func TestUpdateUserAttribute(t *testing.T) {
	now := time.Now().UnixMilli()
	usid := "us123"
	var testcases = []struct {
		name  string
		cred  *cpb.Credential
		attrs []*Attribute
		attr  *Attribute
		out   []*Attribute
	}{
		{name: "empty", cred: &cpb.Credential{Type: cpb.Type_unknown}, attrs: []*Attribute{}, attr: &Attribute{}, out: []*Attribute{}},
		{
			name:  "userbot update userbot",
			cred:  &cpb.Credential{Type: cpb.Type_user, Issuer: usid},
			attrs: []*Attribute{{Key: "fullname", Text: "thanh"}},
			attr:  &Attribute{Key: "fullname", Text: "giang"},
			out:   []*Attribute{{Key: "fullname", ByType: "user", By: usid, UserValue: "giang", Text: "giang", Modified: now}},
		}, {
			name:  "userbot update userbot (no override)",
			attrs: []*Attribute{{Key: "address", Text: "Ha noi"}},
			cred:  &cpb.Credential{Type: cpb.Type_user, Issuer: usid},
			attr:  &Attribute{Key: "address", Text: "Ha tay"},
			out:   []*Attribute{{Key: "address", Text: "Ha noi", UserValue: "Ha tay", Modified: now}}, // address is override prevented
		}, {
			name:  "userbot update userbot empty (no override)",
			attrs: []*Attribute{{Key: "address"}},
			cred:  &cpb.Credential{Type: cpb.Type_user, Issuer: "us99"},
			attr:  &Attribute{Key: "address", Text: "Ha tay"},
			out:   []*Attribute{{Key: "address", By: "us99", ByType: "user", Text: "Ha tay", Modified: now}}, // address is override prevented but is empty => accept write from user
		}, {
			name:  "userbot update agent",
			cred:  &cpb.Credential{Type: cpb.Type_user, Issuer: "us99", ClientId: "bot12"},
			attrs: []*Attribute{{Key: "fullname", ByType: "agent", Text: "thanh"}},
			attr:  &Attribute{Key: "fullname", Text: "giang"},
			out:   []*Attribute{{Key: "fullname", ByType: "agent", Text: "thanh"}},
		}, {
			name:  "userbot update connector",
			cred:  &cpb.Credential{Type: cpb.Type_user, Issuer: "us98"},
			attrs: []*Attribute{{Key: "fullname", ByType: "connector", Text: "thanh"}},
			attr:  &Attribute{Key: "fullname", Text: "giang", ByType: "user", By: "us123"}, // fake user id
			out:   []*Attribute{{Key: "fullname", ByType: "connector", Text: "thanh"}},
		}, {
			name:  "userbot update system 1",
			cred:  &cpb.Credential{Type: cpb.Type_bot, Issuer: "bot", ClientId: "bot"}, // deprecated bot
			attrs: []*Attribute{{Key: "fullname", ByType: "subiz", Text: "thanh"}},
			attr:  &Attribute{Key: "fullname", Text: "giang", ByType: "bot", By: ""},
			out:   []*Attribute{{Key: "fullname", ByType: "subiz", Text: "thanh"}},
		}, {
			name:  "userbot update system 2",
			cred:  &cpb.Credential{Type: cpb.Type_user, Issuer: "us4", ClientId: "bot"},
			attrs: []*Attribute{{Key: "fullname", ByType: "subiz", Text: "thanh"}},
			attr:  &Attribute{Key: "fullname", Text: "giang", ByType: "bot", By: ""},
			out:   []*Attribute{{Key: "fullname", ByType: "subiz", Text: "thanh"}},
		}, {
			name:  "system update on empty by type",
			cred:  &cpb.Credential{Type: cpb.Type_subiz, Issuer: "system"},
			attrs: []*Attribute{{Key: "fullname", Text: "thanh"}},
			attr:  &Attribute{Key: "fullname", Text: "giang", ByType: "subiz", By: ""},
			out:   []*Attribute{{Key: "fullname", ByType: "subiz", Text: "giang", Modified: now, By: "system"}},
		}, {
			name:  "system update on system attribute",
			cred:  &cpb.Credential{Type: cpb.Type_subiz, Issuer: "system"},
			attrs: []*Attribute{{Key: "fullname", ByType: "subiz", Text: "thanh"}},
			attr:  &Attribute{Key: "fullname", Text: "giang", ByType: "subiz", By: ""},
			out:   []*Attribute{{Key: "fullname", ByType: "subiz", Text: "giang", Modified: now, By: "system"}},
		}, {
			name:  "system update on read only",
			cred:  &cpb.Credential{Type: cpb.Type_subiz, Issuer: "system"},
			attrs: []*Attribute{{Key: "created", ByType: "subiz", Datetime: "444"}},
			attr:  &Attribute{Key: "created", Datetime: "3333", ByType: "subiz", By: ""},
			out:   []*Attribute{{Key: "created", ByType: "subiz", Datetime: "3333", Modified: now, By: "system"}},
		}, {
			name:  "system update on agent",
			cred:  &cpb.Credential{Type: cpb.Type_subiz, Issuer: "system"},
			attrs: []*Attribute{{Key: "fullname", ByType: "agent", Text: "thanh"}},
			attr:  &Attribute{Key: "fullname", Text: "giang", ByType: "subiz", By: ""},
			out:   []*Attribute{{Key: "fullname", ByType: "subiz", Text: "giang", Modified: now, By: "system"}},
		}, {
			name:  "update on empty by type",
			cred:  &cpb.Credential{Type: cpb.Type_unknown},
			attrs: []*Attribute{{Key: "fullname", Text: "thanh"}},
			attr:  &Attribute{Key: "fullname", Text: "giang", ByType: "", By: usid},
			out:   []*Attribute{{Key: "fullname", Text: "thanh"}},
		}, {
			name:  "agent update on system ",
			cred:  &cpb.Credential{Type: cpb.Type_agent, Issuer: "ag123"},
			attrs: []*Attribute{{Key: "fullname", ByType: "subiz", Text: "thanh"}},
			attr:  &Attribute{Key: "fullname", Text: "giang", ByType: "agent", By: "ag123"},
			out:   []*Attribute{{Key: "fullname", ByType: "agent", Text: "giang", Modified: now, By: "ag123"}},
		}, {
			name:  "agent update on agent",
			attrs: []*Attribute{{Key: "fullname", ByType: "agent", Text: "thanh"}},
			cred:  &cpb.Credential{Type: cpb.Type_agent, Issuer: "ag123"},
			attr:  &Attribute{Key: "fullname", Text: "giang", ByType: "agent", By: "ag123"},
			out:   []*Attribute{{Key: "fullname", ByType: "agent", Text: "giang", Modified: now, By: "ag123"}},
		}, {
			name:  "agent update on connector",
			attrs: []*Attribute{{Key: "fullname", ByType: "connector", Text: "thanh"}},
			cred:  &cpb.Credential{Type: cpb.Type_agent, Issuer: "ag123"},
			attr:  &Attribute{Key: "fullname", Text: "giang", ByType: "agent", By: "ag123"},
			out:   []*Attribute{{Key: "fullname", ByType: "agent", Text: "giang", Modified: now, By: "ag123"}},
		}, {
			name:  "agent update on system by type readonly",
			cred:  &cpb.Credential{Type: cpb.Type_agent, Issuer: "ag123"},
			attrs: []*Attribute{{Key: "created", Datetime: "123"}},
			attr:  &Attribute{Key: "created", Datetime: "12344", ByType: "agent", By: ""},
			out:   []*Attribute{{Key: "created", Datetime: "123"}},
		}, {
			name:  "connector update on system ",
			cred:  &cpb.Credential{Type: cpb.Type_connector, Issuer: "fabikon", ClientId: "fabikon"},
			attrs: []*Attribute{{Key: "fullname", ByType: "subiz", Text: "thanh"}},
			attr:  &Attribute{Key: "fullname", Text: "giang", ByType: "connector", By: "fabikon"},
			out:   []*Attribute{{Key: "fullname", ByType: "connector", Text: "giang", Modified: now, ConnectorValue: "giang", By: "fabikon"}},
		}, {
			name:  "connector update on agent",
			cred:  &cpb.Credential{Type: cpb.Type_connector, Issuer: "fabikon", ClientId: "fabikon"},
			attrs: []*Attribute{{Key: "fullname", ByType: "agent", Text: "thanh"}},
			attr:  &Attribute{Key: "fullname", Text: "giang", ByType: "connector", By: ""},
			out:   []*Attribute{{Key: "fullname", ByType: "agent", Text: "thanh", ConnectorValue: "giang", Modified: now}},
		},
		{
			name:  "connector update on empty agent",
			cred:  &cpb.Credential{Type: cpb.Type_connector, Issuer: "fabikon", ClientId: "fabikon"},
			attrs: []*Attribute{{Key: "fullname", ByType: "agent", Text: ""}},
			attr:  &Attribute{Key: "fullname", Text: "giang", ByType: "connector", By: ""},
			out:   []*Attribute{{Key: "fullname", ByType: "connector", Text: "giang", ConnectorValue: "giang", Modified: now, By: "fabikon"}},
		}, {
			name:  "connector update on connect",
			cred:  &cpb.Credential{Type: cpb.Type_connector, Issuer: "fabikon", ClientId: "fabikon"},
			attrs: []*Attribute{{Key: "fullname", ByType: "connector", Text: "thanh"}},
			attr:  &Attribute{Key: "fullname", Text: "giang", ByType: "connector", By: "fabikon"},
			out:   []*Attribute{{Key: "fullname", By: "fabikon", ByType: "connector", Text: "giang", ConnectorValue: "giang", Modified: now}},
		}, {
			name:  "connector update on system by type readonly",
			cred:  &cpb.Credential{Type: cpb.Type_connector, Issuer: "fabikon", ClientId: "fabikon"},
			attrs: []*Attribute{{Key: "created", Datetime: "123"}},
			attr:  &Attribute{Key: "created", Datetime: "12344", ByType: "connector", By: "fabikon"},
			out:   []*Attribute{{Key: "created", Datetime: "12344", By: "fabikon", ByType: "connector", ConnectorValue: "12344", Modified: now}},
		}, {
			name:  "system update on system by type",
			cred:  &cpb.Credential{Type: cpb.Type_subiz, Issuer: "system"},
			attrs: []*Attribute{{Key: "created", Datetime: "123"}},
			attr:  &Attribute{Key: "created", Datetime: "12344", ByType: "agent", By: ""},
			out:   []*Attribute{{Key: "created", Datetime: "123"}},
		}, {
			name:  "bot update on system by type",
			cred:  &cpb.Credential{Type: cpb.Type_user, Issuer: usid, ClientId: "bot"},
			attrs: []*Attribute{{Key: "created", Datetime: "123"}},
			attr:  &Attribute{Key: "created", Datetime: "456", ByType: "bot", By: ""},
			out:   []*Attribute{{Key: "created", Datetime: "123"}},
		}, {
			name:  "user update on agent",
			cred:  &cpb.Credential{Type: cpb.Type_user, Issuer: usid}, // deprecated bot
			attrs: []*Attribute{{Key: "fullname", ByType: "agent", Text: "thanh"}},
			attr:  &Attribute{Key: "fullname", Text: "giang", ByType: "user", By: usid},
			out:   []*Attribute{{Key: "fullname", ByType: "agent", Text: "thanh", UserValue: "giang", Modified: now}},
		}, {
			name:  "invalid bot update on empty",
			cred:  &cpb.Credential{Type: cpb.Type_bot, Issuer: "bot", ClientId: "bot"}, // deprecated bot
			attrs: []*Attribute{{Key: "fullname"}},
			attr:  &Attribute{Key: "fullname", Text: "giang", ByType: "bot", By: "bot1"},
			out:   []*Attribute{{Key: "fullname", ByType: "bot", Text: "giang", By: "bot1", Modified: now}},
		}, {
			name:  "invalid bot credential",
			cred:  &cpb.Credential{Type: cpb.Type_bot, Issuer: "bot", ClientId: "bot"}, // deprecated bot
			attrs: []*Attribute{{Key: "fullname", Text: "thanh"}},
			attr:  &Attribute{Key: "fullname", Text: "giang", ByType: "bot", By: "b1"}, // invalid
			out:   []*Attribute{{Key: "fullname", Text: "thanh"}},
		},
		{
			name:  "normal",
			attrs: []*Attribute{{Key: "fullname", Text: "Bui Thi HuongGiang"}},
			cred:  &cpb.Credential{Type: cpb.Type_user, Issuer: usid},
			attr:  &Attribute{Key: "emails", By: usid, ByType: "user", Text: "giangbth@gmail.com"},
			out: []*Attribute{
				{Key: "fullname", Text: "Bui Thi HuongGiang"},
				{
					Key:       "emails",
					Modified:  now,
					ByType:    "user",
					Text:      "giangbth@gmail.com",
					By:        usid,
					UserValue: "giangbth@gmail.com",
				}},
		},
		{
			name:  "user update on primary",
			cred:  &cpb.Credential{Type: cpb.Type_user, Issuer: "us567"},
			attrs: []*Attribute{{Key: "fullname", Text: "Bui Thi Huong Giang", ByType: "user", UserValue: "Bui"}},
			attr:  &Attribute{Key: "fullname", By: "us567", ByType: "user", Text: "Giang"},
			out:   []*Attribute{{Key: "fullname", ByType: "user", Text: "Bui Thi Huong Giang", UserValue: "Bui"}},
		},
		{
			name:  "user push on primary",
			cred:  &cpb.Credential{Type: cpb.Type_user, Issuer: "us567"},
			attrs: []*Attribute{{Key: "fullname", Text: "Bui Thi Huong Giang", ByType: "agent", UserValue: "Bui"}},
			attr:  &Attribute{Key: "fullname", By: "us567", ByType: "user", Text: "Giang", Action: "push"},
			out:   []*Attribute{{Key: "fullname", ByType: "agent", Text: "Bui Thi Huong Giang", UserValue: "Bui"}},
		},
		{
			name:  "update record id",
			cred:  &cpb.Credential{Type: cpb.Type_agent, Issuer: "agrtbnxihehvemfrui"},
			attrs: []*Attribute{{Key: "fullname", Text: "Bui Thi Huong Giang", ByType: "agent", UserValue: "Bui"}},
			attr:  &Attribute{Key: "record_id", Modified: 1701936021404, Text: "abc", ByType: "agent", By: "ag123"},
			out:   []*Attribute{{Key: "fullname", Text: "Bui Thi Huong Giang", ByType: "agent", UserValue: "Bui"}, {Modified: 1701936021404, Key: "record_id", Text: "abc", By: "ag123", ByType: "agent"}},
		},
		{
			name:  "delete specific text value 1",
			cred:  &cpb.Credential{Type: cpb.Type_agent, Issuer: "ag123"},
			attrs: []*Attribute{{Key: "emails", Text: "subiz@gmail.com", OtherValues: []string{"abc@gmail.com"}}},
			attr:  &Attribute{Key: "emails", Text: "subiz@gmail.com", Action: "delete"},
			out:   []*Attribute{{Key: "emails", ByType: "agent", Text: "abc@gmail.com", Modified: now, By: "ag123"}},
		},
		{
			name:  "delete specific text value 2",
			cred:  &cpb.Credential{Type: cpb.Type_agent, Issuer: "ag123"},
			attrs: []*Attribute{{Key: "emails", Text: "subiz@gmail.com", OtherValues: []string{"abc@gmail.com"}}},
			attr:  &Attribute{Key: "emails", Text: "abc@gmail.com", Action: "delete"},
			out:   []*Attribute{{Key: "emails", ByType: "agent", Text: "subiz@gmail.com", Modified: now, By: "ag123"}},
		},
		{
			name:  "delete all text value",
			cred:  &cpb.Credential{Type: cpb.Type_agent, Issuer: "ag123"},
			attrs: []*Attribute{{Key: "emails", ByType: "subiz@gmail.com", Text: "thanh", OtherValues: []string{"abc@gmail.com"}}},
			attr:  &Attribute{Key: "emails", Action: "delete"},
			out:   []*Attribute{{Key: "emails", ByType: "agent", Modified: now, By: "ag123"}},
		},
		{
			name:  "delete specific text value 3",
			cred:  &cpb.Credential{Type: cpb.Type_agent, Issuer: "ag123"},
			attrs: []*Attribute{{Key: "record_id", ByType: "subiz", Text: "aaaaa1702018539570"}, {Key: "emails", ByType: "agent", Text: "aaaaa1702018539570@gmail.com"}},
			attr:  &Attribute{Key: "emails", Text: "aaaaa1702018539570@gmail.com", By: "system", ByType: "subiz", Action: "delete"},
			out:   []*Attribute{{Key: "record_id", ByType: "subiz", Text: "aaaaa1702018539570"}, {Key: "emails", Modified: now, ByType: "subiz", By: "system"}},
		},
	}

	for itc, tc := range testcases {
		// if itc != 34 {
		// continue
		// }
		user := &User{Id: usid, Attributes: tc.attrs}
		cred := tc.cred
		if cred == nil {
			panic("CRED MUST NOT BE EXPLICIT WHEN TEST: " + tc.name)
			// if you just dont care about cred, make it explicit: cred = &cpb.Credential{Type: cpb.Type_unknown}
		}
		UpdateAttribute(cred, DEFS, user, tc.attr, now)
		if diff := isEqualUserAttribute(user.Attributes, tc.out); diff != "" {
			actual, _ := json.Marshal(user.Attributes)
			expected, _ := json.Marshal(tc.out)
			t.Errorf("WRONG TESTCASE #%d: %s\nEXPECT: %s\nACTUAL: %s\nDIFF:%s\n\n", itc+1, tc.name, string(expected), string(actual), diff)
		}
	}
}
