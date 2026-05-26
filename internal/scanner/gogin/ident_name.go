//ff:func feature=scan type=extract control=sequence
//ff:what identName 함수
package gogin

import (
	"go/ast"
)

func identName(expr ast.Expr) string {
	if id, ok := expr.(*ast.Ident); ok {
		return id.Name
	}
	return ""
}
