//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what 선행 대문자 연속(약어)을 소문자화한다 (SMSResult→smsResult, ID→id, Building→building)
package scanner

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func lcFirst(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	// 선행 대문자 연속 개수를 센다
	upper := 0
	for _, r := range runes {
		if !unicode.IsUpper(r) {
			break
		}
		upper++
	}
	if upper == 0 {
		return s
	}
	// 전부 대문자(또는 대문자 1개)면 전부 소문자화
	if upper >= len(runes) {
		return strings.ToLower(s)
	}
	// 대문자 연속이 2개 이상이면, 마지막 대문자는 다음 단어 시작이므로 유지
	if upper == 1 {
		r, size := utf8.DecodeRuneInString(s)
		return string(unicode.ToLower(r)) + s[size:]
	}
	// upper >= 2: SMSResult → smsResult (마지막 대문자 S는 다음 단어 시작 → 유지)
	lowered := make([]rune, len(runes))
	copy(lowered, runes)
	for i := 0; i < upper-1; i++ {
		lowered[i] = unicode.ToLower(runes[i])
	}
	return string(lowered)
}

