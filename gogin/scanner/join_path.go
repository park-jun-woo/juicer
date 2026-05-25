//ff:func feature=scan type=extract control=sequence
//ff:what joinPath 함수
package scanner

import (
	"strings"
)

func joinPath(a, b string) string {
	if a == "" {
		return b
	}
	return strings.TrimRight(a, "/") + "/" + strings.TrimLeft(b, "/")
}
