//ff:func feature=scan type=extract control=selection
//ff:what AST 표현을 사람이 읽을 수 있는 문자열로 변환한다
package fiber

import (
	"fmt"
	"go/ast"
)

func exprString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.SelectorExpr:
		recv := exprString(e.X)
		if recv != "" {
			return recv + "." + e.Sel.Name
		}
		return e.Sel.Name
	case *ast.CompositeLit:
		if e.Type != nil {
			return exprString(e.Type) + "{}"
		}
		return "{}"
	case *ast.StarExpr:
		return "*" + exprString(e.X)
	case *ast.UnaryExpr:
		return exprString(e.X)
	case *ast.CallExpr:
		return exprString(e.Fun) + "()"
	case *ast.IndexExpr:
		return exprString(e.X) + "[" + exprString(e.Index) + "]"
	case *ast.MapType:
		return "map[" + exprString(e.Key) + "]" + exprString(e.Value)
	case *ast.ArrayType:
		return "[]" + exprString(e.Elt)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.BasicLit:
		return e.Value
	}
	return fmt.Sprintf("(%T)", expr)
}
