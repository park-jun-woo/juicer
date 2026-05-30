//ff:func feature=scan type=test control=sequence
//ff:what typeOfVar 테스트 헬퍼
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

// typeOfVar type-checks src and returns the type of the named package-level var.
func typeOfVar(t *testing.T, src, varName string) types.Type {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{Defs: map[*ast.Ident]types.Object{}}
	pkg, err := conf.Check("m", fset, []*ast.File{file}, info)
	if err != nil {
		t.Fatal(err)
	}
	obj := pkg.Scope().Lookup(varName)
	if obj == nil {
		t.Fatalf("var %s not found", varName)
	}
	return obj.Type()
}
