//ff:func feature=scan type=test control=iteration dimension=1
//ff:what handlerBody 테스트 헬퍼
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func handlerBody(t *testing.T, src string) *ast.BlockStmt {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Body != nil {
			return fn.Body
		}
	}
	t.Fatal("no func body")
	return nil
}
