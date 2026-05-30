//ff:func feature=scan type=test control=sequence
//ff:what routeCallFrom 테스트 헬퍼
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func routeCallFrom(t *testing.T, src string) (*ast.CallExpr, *token.FileSet) {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parserParseFile(fset, src)
	if err != nil {
		t.Fatal(err)
	}
	var call *ast.CallExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok && call == nil {
			if _, ok := c.Fun.(*ast.SelectorExpr); ok {
				call = c
			}
		}
		return true
	})
	return call, fset
}
