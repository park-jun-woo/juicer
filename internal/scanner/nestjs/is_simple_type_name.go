//ff:func feature=scan type=convert control=iteration dimension=1 topic=nestjs
//ff:what 문자열이 컴포넌트 스키마명이 될 수 있는 단순 식별자인지 검사한다
package nestjs

// isSimpleTypeName reports whether s is a bare identifier (no generics, unions,
// arrays, members, whitespace, etc.) that can name a component schema.
func isSimpleTypeName(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !(r == '_' || r == '$' ||
			(r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9')) {
			return false
		}
	}
	return true
}
