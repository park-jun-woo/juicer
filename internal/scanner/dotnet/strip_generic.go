//ff:func feature=scan type=convert control=sequence topic=dotnet
//ff:what 제네릭 타입에서 베이스 이름만 추출한다
package dotnet

import "strings"

func stripGeneric(t string) string {
	idx := strings.Index(t, "<")
	if idx > 0 {
		return t[:idx]
	}
	return t
}
