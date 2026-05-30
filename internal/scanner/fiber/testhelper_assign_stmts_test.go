//ff:func feature=scan type=test control=sequence
//ff:what assignStmts 테스트 헬퍼
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func assignStmts(t *testing.T, src string) []*ast.AssignStmt {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parserParseFile(fset, src)
	if err != nil {
		t.Fatal(err)
	}
	var out []*ast.AssignStmt
	ast.Inspect(file, func(n ast.Node) bool {
		if a, ok := n.(*ast.AssignStmt); ok {
			out = append(out, a)
		}
		return true
	})
	return out
}
