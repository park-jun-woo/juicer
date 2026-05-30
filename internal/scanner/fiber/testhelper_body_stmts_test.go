//ff:func feature=scan type=test control=iteration dimension=1
//ff:what bodyStmts 테스트 헬퍼
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func bodyStmts(t *testing.T, src string) []ast.Stmt {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Body != nil {
			return fn.Body.List
		}
	}
	t.Fatal("no func body")
	return nil
}
