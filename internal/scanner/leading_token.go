//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what operationId의 선두 토큰(첫 camelCase/snake 단어)을 소문자로 추출한다
package scanner

import (
	"strings"
	"unicode"
)

func leadingToken(id string) string {
	if id == "" {
		return ""
	}
	runes := []rune(id)
	end := len(runes)
	for i := 1; i < len(runes); i++ {
		if runes[i] == '_' || unicode.IsUpper(runes[i]) {
			end = i
			break
		}
	}
	return strings.ToLower(string(runes[:end]))
}
