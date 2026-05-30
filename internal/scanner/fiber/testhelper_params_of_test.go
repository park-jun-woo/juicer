//ff:func feature=scan type=test control=iteration dimension=1
//ff:what paramsOf 테스트 헬퍼
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func paramsOf(t *testing.T, src string) *ast.FieldList {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok {
			return fn.Type.Params
		}
	}
	t.Fatal("no func")
	return nil
}
