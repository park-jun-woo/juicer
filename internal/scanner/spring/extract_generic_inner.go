//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 제네릭 타입에서 내부 타입을 추출한다
package spring

import "strings"

func extractGenericInner(jtype string) string {
	idx := strings.Index(jtype, "<")
	if idx < 0 {
		return ""
	}
	inner := jtype[idx+1 : len(jtype)-1]
	return strings.TrimSpace(inner)
}
