//ff:func feature=scan type=extract control=sequence
//ff:what AST 타입 표현이 *fiber.Ctx인지 검사한다
package fiber

import (
	"go/ast"
)

func isFiberContextType(expr ast.Expr) bool {
	star, ok := expr.(*ast.StarExpr)
	if !ok {
		return false
	}
	sel, ok := star.X.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	id, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}
	return sel.Sel.Name == "Ctx" && id.Name == "fiber"
}
