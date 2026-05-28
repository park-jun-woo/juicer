//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 제네릭 타입에서 타입 인자를 추출한다
package quarkus

import "strings"

func extractGenericArgs(typeName string) string {
	idx := strings.Index(typeName, "<")
	if idx < 0 {
		return ""
	}
	return strings.TrimSpace(typeName[idx+1 : len(typeName)-1])
}
