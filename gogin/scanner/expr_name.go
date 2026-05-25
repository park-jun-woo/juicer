//ff:func feature=scan type=extract control=selection
//ff:what exprName 함수
package scanner

import (
	"go/ast"
)

func exprName(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.SelectorExpr:
		if recv := identName(e.X); recv != "" {
			return recv + "." + e.Sel.Name
		}
		return e.Sel.Name
	case *ast.FuncLit:
		return "(inline)"
	case *ast.CallExpr:
		return exprName(e.Fun) + "()"
	}
	return ""
}
