//ff:func feature=scan type=test control=sequence
//ff:what firstFuncLit 테스트 헬퍼
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

// firstFuncLit parses a Go source and returns the first FuncLit found.
func firstFuncLit(t *testing.T, src string) *ast.FuncLit {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	var lit *ast.FuncLit
	ast.Inspect(file, func(n ast.Node) bool {
		if fl, ok := n.(*ast.FuncLit); ok && lit == nil {
			lit = fl
			return false
		}
		return true
	})
	if lit == nil {
		t.Fatal("no FuncLit found")
	}
	return lit
}
