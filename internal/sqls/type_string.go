//ff:func feature=sql type=parse control=selection
//ff:what ast.Expr를 Go 타입 문자열로 렌더링
package sqls

import (
	"fmt"
	"go/ast"
)

// typeString renders an ast.Expr as a Go type string.
//
func typeString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + typeString(t.X)
	case *ast.SelectorExpr:
		return typeString(t.X) + "." + t.Sel.Name
	case *ast.ArrayType:
		return "[]" + typeString(t.Elt)
	case *ast.MapType:
		return "map[" + typeString(t.Key) + "]" + typeString(t.Value)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.Ellipsis:
		return "..." + typeString(t.Elt)
	default:
		return fmt.Sprintf("%T", expr)
	}
}

