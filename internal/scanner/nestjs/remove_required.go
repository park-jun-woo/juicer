//ff:func feature=scan type=convert control=iteration dimension=1 topic=nestjs
//ff:what validate 문자열에서 "required" 토큰을 제거한다
package nestjs

import "strings"

// removeRequired removes the "required" token from a validate string.
func removeRequired(validate string) string {
	parts := strings.Split(validate, ",")
	var out []string
	for _, p := range parts {
		if p != "required" {
			out = append(out, p)
		}
	}
	return strings.Join(out, ",")
}
