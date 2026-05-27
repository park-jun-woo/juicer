//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what 엔드포인트 목록의 각 핸들러 body를 분석하여 요청/응답 정보를 채운다
package gogin

import (
	"go/ast"

	"golang.org/x/tools/go/packages"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func analyzeHandlers(pkgs []*packages.Package, endpoints []scanner.Endpoint, root string, handlerExprsMap map[int][]ast.Expr, idx *funcIndex) {
	for i := range endpoints {
		ep := &endpoints[i]
		exprs := handlerExprsMap[i]
		if len(exprs) == 0 {
			continue
		}

		// 각 핸들러 표현(핸들러 + 미들웨어 핸들러)에서 가장 마지막이 실제 핸들러
		for _, expr := range exprs {
			info := findInfoForExpr(expr, pkgs)
			if info == nil {
				continue
			}
			analyzeExpr(ep, expr, info, idx)
		}
	}
}
