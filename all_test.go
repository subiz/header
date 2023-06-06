package header

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"testing"

	pb "github.com/subiz/header/account"
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
		{",;84123123123,", "0123123123"},
	}

	for i, tc := range testcases {
		if tc.out != NormPhone(tc.in) {
			t.Errorf("WRONG AT TEST #%d, expect %s, got %s", i+1, tc.out, NormPhone(tc.in))
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
	shardNumber := int(crc32.ChecksumIEEE([]byte("acriviayfmabzskstrpq"))) % 4
	fmt.Println(shardNumber)
}

func TestUnpack(t *testing.T) {
	str := "1a196576726977616f686f646a686e787a6f71776668686f67757940b7dde7bc94304a0e636f6e74656e745f766965776564a2013f323d3a3b68747470733a2f2f6170702e737562697a2e636f6d2e766e2f636f6e766f3f7569643d2d266369643d63737269767a71716b706b666b726d726c78920353121a1a184f4f697a56576f51554570533652497444704d2b32413d3d1a297573616371717665626a69667768626c7763616762665f6167717176666a6e6a647177726778706863220a636c737562697a617069"

	bs, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}

	ev := &Event{}
	proto.Unmarshal(bs, ev)
	out, _ := json.Marshal(ev)
	fmt.Println(string(out))
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

func TestUpdateUser(t *testing.T) {
	user := &User{}
	UpdateAttr(user, "fullname", "string", "thanh", "upsert")
	UpdateAttr(user, "fullname", "string", "dao", "upsert")
	UpdateAttr(user, "fullname", "stringd", "thanh", "upsert") // invalid
	b, _ := json.Marshal(user)
	fmt.Println(string(b))
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
