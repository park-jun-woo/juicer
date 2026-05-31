//ff:func feature=scan type=extract control=selection topic=express
//ff:what res 호출이 없는 핸들러의 메서드별 관례 기본 응답을 반환한다 (POST→201, DELETE→204, 그 외 200)
package express

import (
	"strings"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func defaultResponseForMethod(method string) []scanner.Response {
	switch strings.ToUpper(method) {
	case "POST":
		return []scanner.Response{{Status: "201", Kind: "json"}}
	case "DELETE":
		return []scanner.Response{{Status: "204", Kind: "json"}}
	default:
		return []scanner.Response{{Status: "200", Kind: "json"}}
	}
}
