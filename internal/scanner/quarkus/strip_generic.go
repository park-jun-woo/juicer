//ff:func feature=scan type=extract control=sequence topic=quarkus
//ff:what 제네릭 타입에서 래퍼 이름을 분리한다
package quarkus

import "strings"

func stripGeneric(typeName string) string {
	idx := strings.Index(typeName, "<")
	if idx < 0 {
		return typeName
	}
	return typeName[:idx]
}
