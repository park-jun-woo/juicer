//ff:func feature=scan type=convert control=iteration dimension=1 topic=nestjs
//ff:what 문자열에서 제네릭 타입 파라미터를 실제 타입으로 단어 경계 기반 치환한다
package nestjs

import "regexp"

// substituteOne replaces all occurrences of generic type parameter names in s
// with their mapped concrete types, using word-boundary matching.
func substituteOne(s string, typeMap map[string]string) string {
	if s == "" {
		return s
	}
	for old, new_ := range typeMap {
		re := regexp.MustCompile(`\b` + regexp.QuoteMeta(old) + `\b`)
		s = re.ReplaceAllString(s, new_)
	}
	return s
}
