//ff:func feature=scan type=convert control=selection topic=nestjs
//ff:what TS 배열 래퍼(T[] / Array<T>)를 벗기고 요소 타입과 배열 여부를 반환한다
package nestjs

import "strings"

// stripArrayWrapper removes a TS array wrapper ("T[]" or "Array<T>") and returns
// the element type plus whether a wrapper was present.
func stripArrayWrapper(ts string) (elem string, isArray bool) {
	switch {
	case strings.HasSuffix(ts, "[]"):
		return strings.TrimSpace(ts[:len(ts)-2]), true
	case strings.HasPrefix(ts, "Array<") && strings.HasSuffix(ts, ">"):
		return strings.TrimSpace(ts[6 : len(ts)-1]), true
	}
	return ts, false
}
