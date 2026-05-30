//ff:func feature=scan type=test control=iteration dimension=1
//ff:what funcDeclFrom 테스트 헬퍼
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func funcDeclFrom(t *testing.T, src string) *ast.FuncDecl {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok {
			return fn
		}
	}
	t.Fatal("no func decl")
	return nil
}
