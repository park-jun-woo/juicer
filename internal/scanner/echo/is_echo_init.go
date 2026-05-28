//ff:func feature=scan type=extract control=sequence
//ff:what echo.New() 호출을 감지한다
package echo

import (
	"go/ast"
)

func isEchoInit(sel *ast.SelectorExpr, alias string) bool {
	id, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}
	return id.Name == alias && sel.Sel.Name == "New"
}
