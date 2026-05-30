//ff:func feature=scan type=test control=sequence
//ff:what structFields 테스트 헬퍼
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

// structFields type-checks src and returns the named struct's fields as *types.Var.
func structFields(t *testing.T, src, typeName string) (*types.Struct, *types.Info) {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Defs:  map[*ast.Ident]types.Object{},
		Types: map[ast.Expr]types.TypeAndValue{},
	}
	pkg, err := conf.Check("m", fset, []*ast.File{file}, info)
	if err != nil {
		t.Fatal(err)
	}
	obj := pkg.Scope().Lookup(typeName)
	if obj == nil {
		t.Fatalf("type %s not found", typeName)
	}
	st, ok := obj.Type().Underlying().(*types.Struct)
	if !ok {
		t.Fatalf("%s is not a struct", typeName)
	}
	return st, info
}
