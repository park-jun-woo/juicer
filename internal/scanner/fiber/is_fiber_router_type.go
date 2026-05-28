//ff:func feature=scan type=extract control=sequence
//ff:what AST 타입 표현이 *fiber.App 또는 *fiber.Group인지 검사한다
package fiber

import (
	"go/ast"
)

func isFiberRouterType(expr ast.Expr, alias string) bool {
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
	return id.Name == alias && fiberRouterTypes[sel.Sel.Name]
}
