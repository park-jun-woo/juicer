//ff:func feature=scan type=extract control=sequence
//ff:what isGinRouterType 함수
package gogin

import (
	"go/ast"
)

func isGinRouterType(expr ast.Expr, alias string) bool {
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
	return id.Name == alias && ginRouterTypes[sel.Sel.Name]
}
