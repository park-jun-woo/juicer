//ff:func feature=scan type=convert control=sequence topic=fastapi
//ff:what nullable 래퍼를 제거하여 내부 타입을 반환한다
package fastapi

import "strings"

// unwrapNullable removes the nullable wrapper from a type.
func unwrapNullable(typeName string) string {
	if strings.HasPrefix(typeName, "Optional[") && strings.HasSuffix(typeName, "]") {
		return strings.TrimSpace(typeName[9 : len(typeName)-1])
	}
	if strings.Contains(typeName, "|") {
		parts := splitTopLevel(typeName)
		nonNone := filterNone(parts)
		if len(nonNone) == 1 {
			return nonNone[0]
		}
	}
	return typeName
}
