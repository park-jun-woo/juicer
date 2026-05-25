//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what pathParams 함수
package scanner

import (
	"strings"
)

func pathParams(path string) []Param {
	var params []Param
	for _, seg := range strings.Split(path, "/") {
		if strings.HasPrefix(seg, ":") {
			params = append(params, Param{Name: seg[1:], Type: "string"})
		}
	}
	return params
}
