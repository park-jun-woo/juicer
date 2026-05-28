//ff:func feature=scan type=extract control=sequence
//ff:what 바인딩 인자(&req)에서 변수명을 추출한다
package fiber

import (
	"go/ast"
	"go/token"
)

func bindVarName(expr ast.Expr) string {
	if unary, ok := expr.(*ast.UnaryExpr); ok && unary.Op == token.AND {
		return exprString(unary.X)
	}
	return exprString(expr)
}
