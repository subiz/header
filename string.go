package header

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// take first n runes of the string
// the-original-string
// 01234567890
// 00111110000
//
//	[    )
func FirstN(s string, end int) string {
	if end <= 0 {
		return ""
	}

	if s == "" {
		return ""
	}

	if end >= len(s) {
		return s
	}

	i := 0
	for j := range s {
		if i == end {
			return s[:j]
		}
		i++
	}
	return s
}

var VNMAP = map[rune]rune{
	'ạ': 'a', 'ả': 'a', 'ã': 'a', 'à': 'a', 'á': 'a', 'â': 'a', 'ậ': 'a', 'ầ': 'a', 'ấ': 'a',
	'ẩ': 'a', 'ẫ': 'a', 'ă': 'a', 'ắ': 'a', 'ằ': 'a', 'ặ': 'a', 'ẳ': 'a', 'ẵ': 'a',
	'ó': 'o', 'ò': 'o', 'ọ': 'o', 'õ': 'o', 'ỏ': 'o', 'ô': 'o', 'ộ': 'o', 'ổ': 'o', 'ỗ': 'o',
	'ồ': 'o', 'ố': 'o', 'ơ': 'o', 'ờ': 'o', 'ớ': 'o', 'ợ': 'o', 'ở': 'o', 'ỡ': 'o',
	'é': 'e', 'è': 'e', 'ẻ': 'e', 'ẹ': 'e', 'ẽ': 'e', 'ê': 'e', 'ế': 'e', 'ề': 'e', 'ệ': 'e', 'ể': 'e', 'ễ': 'e',
	'ú': 'u', 'ù': 'u', 'ụ': 'u', 'ủ': 'u', 'ũ': 'u', 'ư': 'u', 'ự': 'u', 'ữ': 'u', 'ử': 'u', 'ừ': 'u', 'ứ': 'u',
	'í': 'i', 'ì': 'i', 'ị': 'i', 'ỉ': 'i', 'ĩ': 'i',
	'ý': 'y', 'ỳ': 'y', 'ỷ': 'y', 'ỵ': 'y', 'ỹ': 'y',
	'đ': 'd',
	'Ạ': 'A', 'Ả': 'A', 'Ã': 'A', 'À': 'A', 'Á': 'A', 'Â': 'A', 'Ậ': 'A', 'Ầ': 'A', 'Ấ': 'A',
	'Ẩ': 'A', 'Ẫ': 'A', 'Ă': 'A', 'Ắ': 'A', 'Ằ': 'A', 'Ặ': 'A', 'Ẳ': 'A', 'Ẵ': 'A',
	'Ó': 'O', 'Ò': 'O', 'Ọ': 'O', 'Õ': 'O', 'Ỏ': 'O', 'Ô': 'O', 'Ộ': 'O', 'Ổ': 'O', 'Ỗ': 'O',
	'Ồ': 'O', 'Ố': 'O', 'Ơ': 'O', 'Ờ': 'O', 'Ớ': 'O', 'Ợ': 'O', 'Ở': 'O', 'Ỡ': 'O',
	'É': 'E', 'È': 'E', 'Ẻ': 'E', 'Ẹ': 'E', 'Ẽ': 'E', 'Ê': 'E', 'Ế': 'E', 'Ề': 'E', 'Ệ': 'E', 'Ể': 'E', 'Ễ': 'E',
	'Ú': 'U', 'Ù': 'U', 'Ụ': 'U', 'Ủ': 'U', 'Ũ': 'U', 'Ư': 'U', 'Ự': 'U', 'Ữ': 'U', 'Ử': 'U', 'Ừ': 'U', 'Ứ': 'U',
	'Í': 'I', 'Ì': 'I', 'Ị': 'I', 'Ỉ': 'I', 'Ĩ': 'I',
	'Ý': 'Y', 'Ỳ': 'Y', 'Ỷ': 'Y', 'Ỵ': 'Y', 'Ỹ': 'Y',
	'Đ': 'D',
}

// Ascii replaces all non-ascii characters to equivalent ascii characters
// e.g: â => a, đ => d, ...
// To ensure the output string is pure ascii, this function remove all
// characters that does not have equivalent ascii character, for example: 主
func Ascii(text string) string {
	return strings.Map(func(r rune) rune {
		if r <= unicode.MaxASCII {
			return r
		}

		// fast case: only ascii or vietnamese
		if vnr := VNMAP[r]; vnr != 0 {
			return vnr
		}

		// remove all non-ascii
		return -1
	}, text)
}

// Norm trims spaces, removes invalid Unicode runes, and safely truncates
// the string to at most `max` runes (without splitting UTF-8 characters).
func Norm(s string, maxChar int) string {
	out := make([]rune, 0, len(s))
	for _, r := range s {
		if !isInvalidRune(r) {
			out = append(out, r)
		}
	}
	s = strings.TrimSpace(string(out))
	if s == "" {
		return ""
	}

	if maxChar <= 0 {
		maxChar = len(s)
	}

	runeCount := utf8.RuneCountInString(s)
	if maxChar >= runeCount {
		return strings.TrimSpace(s)
	}
	i := 0
	for j := range s {
		if i == maxChar {
			return strings.TrimSpace(s[:j])
		}
		i++
	}
	return strings.TrimSpace(s)
}

// isInvalidRune reports whether a rune is invalid or unwanted for text display.
func isInvalidRune(r rune) bool {
	// Invalid UTF-8 placeholder
	if r == utf8.RuneError {
		return true
	}

	// C0 and C1 control characters, except \n and \t
	if (r < 0x20 && r != '\n' && r != '\t') || (r >= 0x7f && r < 0xa0) {
		return true
	}

	// Unicode noncharacters: U+FDD0–U+FDEF or any codepoint ending in FFFE/FFFF
	if (r >= 0xFDD0 && r <= 0xFDEF) || (r&0xFFFF == 0xFFFE) || (r&0xFFFF == 0xFFFF) {
		return true
	}

	// Optionally remove unprintable runes (if you want only visible text)
	if !unicode.IsPrint(r) && !unicode.IsSpace(r) {
		return true
	}

	return false
}
