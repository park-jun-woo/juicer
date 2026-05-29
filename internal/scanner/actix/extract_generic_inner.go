//ff:func feature=scan type=convert control=sequence topic=actix
//ff:what 제네릭 타입 문자열에서 첫 꺾쇠 안쪽 타입을 추출한다
package actix

import "strings"

func extractGenericInner(t string) string {
	idx := strings.Index(t, "<")
	if idx < 0 {
		return t
	}
	inner := t[idx+1:]
	if len(inner) > 0 && inner[len(inner)-1] == '>' {
		inner = inner[:len(inner)-1]
	}
	return strings.TrimSpace(inner)
}
