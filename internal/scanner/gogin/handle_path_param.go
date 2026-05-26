//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what c.Param("name") 호출에서 경로 파라미터를 추출한다 (Phase 001 보완)
package gogin

import (
	"go/ast"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func handlePathParam(ep *scanner.Endpoint, call *ast.CallExpr) {
	if len(call.Args) < 1 {
		return
	}
	name := stringLitValue(call.Args[0])
	if name == "" {
		return
	}
	scanner.EnsureRequest(ep)

	// Phase 001에서 경로 패턴으로 이미 추출했을 수 있음 — 중복 방지
	for _, p := range ep.Request.PathParams {
		if p.Name == name {
			return
		}
	}
	ep.Request.PathParams = append(ep.Request.PathParams, scanner.Param{Name: name, Type: "string"})
}

