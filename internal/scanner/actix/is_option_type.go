//ff:func feature=scan type=convert control=sequence topic=actix
//ff:what 타입이 Option<...> 인지 판별한다
package actix

import "strings"

func isOptionType(t string) bool {
	return strings.HasPrefix(t, "Option<")
}
