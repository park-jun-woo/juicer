//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what 엔드포인트 목록의 각 핸들러 body를 분석하여 요청/응답 정보를 채운다
package scanner

import (
	"golang.org/x/tools/go/packages"
)

func analyzeHandlers(pkgs []*packages.Package, endpoints []Endpoint, root string) {
	idx := buildFuncIndex(pkgs)

	for i := range endpoints {
		ep := &endpoints[i]
		if len(ep.handlerExprs) == 0 {
			continue
		}

		// 각 핸들러 표현(핸들러 + 미들웨어 핸들러)에서 가장 마지막이 실제 핸들러
		for _, expr := range ep.handlerExprs {
			info := findInfoForExpr(expr, pkgs)
			if info == nil {
				continue
			}
			analyzeExpr(ep, expr, info, idx)
		}

		// handlerExprs 정리 (직렬화에 포함 안 되지만 메모리 해제)
		ep.handlerExprs = nil
	}
}

