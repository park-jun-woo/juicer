//ff:func feature=scan type=extract control=sequence
//ff:what isGinInit 함수
package scanner

import (
	"go/ast"
)

func isGinInit(sel *ast.SelectorExpr, alias string) bool {
	id, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}
	return id.Name == alias && (sel.Sel.Name == "Default" || sel.Sel.Name == "New")
}
