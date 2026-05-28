//ff:func feature=scan type=extract control=sequence
//ff:what unquote 함수
package echo

import (
	"strconv"
	"strings"
)

func unquote(s string) string {
	if u, err := strconv.Unquote(s); err == nil {
		return u
	}
	return strings.Trim(s, `"`+"`")
}
