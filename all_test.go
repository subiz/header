package header

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"

	"testing"
)

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

func TestDeltaToPlainText(t *testing.T) {
	out := DeltaToPlainText(`{"ops":[{"insert":"xin choa "},{"insert":{"mention":{"fullname":"Trang Nguyễn","id":"agqnjeturduiovlkyn"}}},{"insert":" day la mot emoji "},{"insert":{"emoji":"neutral"}},{"insert":"\n"}]}`)
	println(out)
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
