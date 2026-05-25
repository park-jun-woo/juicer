//ff:func feature=scan type=extract control=sequence
//ff:what 문자열 리터럴 AST 노드에서 값을 추출한다
package scanner

import (
	"go/ast"
	"go/token"
)

func stringLitValue(expr ast.Expr) string {
	lit, ok := expr.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return ""
	}
	return unquote(lit.Value)
}

