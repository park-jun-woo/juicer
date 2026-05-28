//ff:func feature=scan type=extract control=selection
//ff:what 리터럴 또는 문자열 연결에서 경로를 추출한다
package fiber

import (
	"go/ast"
	"go/token"
)

// extractPathString extracts a path from a literal or string concatenation.
func extractPathString(expr ast.Expr) (string, bool) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			return unquote(e.Value), true
		}
	case *ast.BinaryExpr:
		return extractBinaryPath(e)
	}
	return "", false
}
