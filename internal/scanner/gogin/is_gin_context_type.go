//ff:func feature=scan type=extract control=sequence
//ff:what AST 타입 표현이 *gin.Context인지 검사한다
package gogin

import (
	"go/ast"
)

func isGinContextType(expr ast.Expr) bool {
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
	return sel.Sel.Name == "Context" && id.Name == "gin"
}

