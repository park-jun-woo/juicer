//ff:func feature=scan type=convert control=sequence topic=nestjs
//ff:what Promise<X> 래핑을 제거한다
package nestjs

import "strings"

// unwrapPromise removes Promise<...> wrapper if present.
func unwrapPromise(t string) string {
	t = strings.TrimSpace(t)
	if strings.HasPrefix(t, "Promise<") && strings.HasSuffix(t, ">") {
		return strings.TrimSpace(t[8 : len(t)-1])
	}
	return t
}
