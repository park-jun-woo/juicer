//ff:func feature=scan type=extract control=sequence
//ff:what AST 타입 표현이 gin.H SelectorExpr인지 검사한다
package scanner

import "go/ast"

// isGinHSelector checks if an AST type expression is gin.H selector.
func isGinHSelector(typeExpr ast.Expr) bool {
	sel, ok := typeExpr.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	id, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}
	return id.Name == "gin" && sel.Sel.Name == "H"
}
