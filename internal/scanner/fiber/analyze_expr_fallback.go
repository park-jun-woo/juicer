//ff:func feature=scan type=extract control=sequence
//ff:what 타입정보가 없을 때 named-func 핸들러를 AST로 찾아 body를 폴백 분석한다
package fiber

import (
	"go/ast"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func analyzeExprFallback(ep *scanner.Endpoint, expr ast.Expr, idx *funcIndex) {
	if lit, ok := expr.(*ast.FuncLit); ok {
		scanFallbackFunc(ep, lit.Type, lit.Body, idx)
		return
	}
	name := handlerFuncName(expr)
	if name == "" {
		return
	}
	fn := idx.byName[name]
	if fn == nil {
		return
	}
	scanFallbackFunc(ep, fn.Type, fn.Body, idx)
}
