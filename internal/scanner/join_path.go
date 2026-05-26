//ff:func feature=scan type=extract control=sequence
//ff:what JoinPath 함수
package scanner

import (
	"strings"
)

func JoinPath(a, b string) string {
	if a == "" {
		return b
	}
	return strings.TrimRight(a, "/") + "/" + strings.TrimLeft(b, "/")
}
