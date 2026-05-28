//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what Actix-web 경로에서 {param} 형식의 경로 파라미터를 추출한다
package actix

import (
	"strings"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func extractPathParams(path string) []scanner.Param {
	var params []scanner.Param
	for _, seg := range strings.Split(path, "/") {
		if strings.HasPrefix(seg, "{") && strings.HasSuffix(seg, "}") {
			name := seg[1 : len(seg)-1]
			params = append(params, scanner.Param{Name: name, Type: "string"})
		}
	}
	return params
}
