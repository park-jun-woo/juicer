//ff:func feature=scan type=extract control=sequence
//ff:what fiber.New() 호출을 감지한다
package fiber

import (
	"go/ast"
)

func isFiberInit(sel *ast.SelectorExpr, alias string) bool {
	id, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}
	return id.Name == alias && sel.Sel.Name == "New"
}
