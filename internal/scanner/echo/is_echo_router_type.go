//ff:func feature=scan type=extract control=sequence
//ff:what AST 타입이 *echo.Echo 또는 *echo.Group인지 검사한다
package echo

import (
	"go/ast"
)

func isEchoRouterType(expr ast.Expr, alias string) bool {
	if star, ok := expr.(*ast.StarExpr); ok {
		expr = star.X
	}
	sel, ok := expr.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	id, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}
	return id.Name == alias && echoRouterTypes[sel.Sel.Name]
}
