//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 문자열이 dotted identifier 형식인지 판별한다
package fastapi

import "strings"

// isDottedIdentifier returns true if s has the form "ident.ident" (exactly one
// dot separating two valid Python identifiers).
func isDottedIdentifier(s string) bool {
	parts := strings.SplitN(s, ".", 2)
	if len(parts) != 2 {
		return false
	}
	return isIdentifier(parts[0]) && isIdentifier(parts[1])
}
