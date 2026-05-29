//ff:func feature=prisma type=convert control=sequence topic=prisma
//ff:what 비어있지 않은 문자열만 슬라이스에 추가
package prisma

// appendIfNotEmpty appends s to dst only when s is non-empty.
func appendIfNotEmpty(dst []string, s string) []string {
	if s == "" {
		return dst
	}
	return append(dst, s)
}
