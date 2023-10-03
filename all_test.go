package header

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"strings"
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
	str := "0x12126373727563616f7a6a6d626a62787a6b6e611a14616372737a666d6172766f6972676474677467773a4c1a047573657222157573727563616f7a63776e67617161797a616e756e62086f627365727665728201121a1031363937353132343037313738383738a001f6fdc8a1ae31a801f6fdc8a1ae316a301a1132333839393738373339323939383136322a0866616365626f6f6bba0110313639373531323430373137383837389801c5a4e69cae31b00197e6dda806ba0193031a1965767275636a7a656665627569706b6863646f75666f6879642214616372737a666d6172766f69726764746774677740c5a4e69cae314a0c6d6573736167655f73656e74a201a4021aa1021a09706c61696e74657874229701120a696d6167652f6a7065671a7568747470733a2f2f7663646e2e737562697a2d63646e2e636f6d2f66696c652f376666643765396266303535653136633332653037373132316364623831613932326265313131323531393132643130336134363164333632653732336639385f616372737a666d6172766f6972676474677467772a08756e7469746c6564620466696c6568a9b814526612085f666162696b6f6e1a5a226d5f373445445a4858367346553673425f475647614752533771764356566c6e31455f52364b495a5a63677676524244675f36364d6b7058704e6f6d515a4a43653463693043346572436677794a665848654767444e4c772272126373727563616f7a6a6d626a62787a6b6e619203221a10313639373531323430373137383837382207666162696b6f6e320564756d6d79d80298f8b2a3ae31"
	str = strings.TrimPrefix(str, "0x")

	bs, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}

	ev := &Conversation{}
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
