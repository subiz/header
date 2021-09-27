package header

import (
	"fmt"
	"testing"
)

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

	if "tieng anh" != GetI18n(str, "en-US", "ja-JP") {
		t.Fatalf("should be tieng anh, got %s", GetI18n(str, "en-US", "jp-JP"))
	}

	if "tieng nhat" != GetI18n(str, "ar-MA", "ja-JP") {
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
