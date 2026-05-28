//ff:func feature=scan type=extract control=sequence
//ff:what AST 타입 표현이 echo.Context인지 검사한다
package echo

import (
	"go/ast"
)

// isEchoContextType checks if an AST type expression is echo.Context.
// Note: Echo uses an interface (echo.Context), not a pointer (*gin.Context).
func isEchoContextType(expr ast.Expr) bool {
	sel, ok := expr.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	id, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}
	return sel.Sel.Name == "Context" && id.Name == "echo"
}
