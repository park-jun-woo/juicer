//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what c.Query("name") 호출에서 쿼리 파라미터 이름을 추출한다
package fiber

import (
	"go/ast"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func handleQuery(ep *scanner.Endpoint, call *ast.CallExpr, method string) {
	if len(call.Args) < 1 {
		return
	}
	name := stringLitValue(call.Args[0])
	if name == "" {
		return
	}
	scanner.EnsureRequest(ep)

	// 중복 방지
	for _, q := range ep.Request.Query {
		if q.Name == name {
			return
		}
	}

	p := scanner.Param{Name: name, Type: "string"}
	// Fiber's c.Query("name", "default") has optional second arg as default
	if len(call.Args) >= 2 {
		p.Default = stringLitValue(call.Args[1])
	}
	ep.Request.Query = append(ep.Request.Query, p)
}
