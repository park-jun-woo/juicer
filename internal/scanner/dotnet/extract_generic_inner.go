//ff:func feature=scan type=convert control=sequence topic=dotnet
//ff:what 제네릭 타입에서 내부 타입을 추출한다
package dotnet

import "strings"

func extractGenericInner(t string) string {
	start := strings.Index(t, "<")
	if start < 0 {
		return t
	}
	end := strings.LastIndex(t, ">")
	if end < 0 || end <= start {
		return t
	}
	return strings.TrimSpace(t[start+1 : end])
}
