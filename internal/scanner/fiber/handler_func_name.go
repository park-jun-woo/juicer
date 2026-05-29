//ff:func feature=scan type=extract control=selection
//ff:what 핸들러 AST 표현(Ident/SelectorExpr)에서 함수 이름을 추출한다
package fiber

import (
	"go/ast"
)

func handlerFuncName(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.SelectorExpr:
		return e.Sel.Name
	case *ast.CallExpr:
		return handlerFuncName(e.Fun)
	}
	return ""
}
