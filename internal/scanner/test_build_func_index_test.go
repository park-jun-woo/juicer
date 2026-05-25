//ff:func feature=scan type=extract control=sequence
//ff:what TestBuildFuncIndex 테스트
package scanner

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestBuildFuncIndex(t *testing.T) {
	src := `package test
func Hello() {}
func World() {}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "test.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}

	conf := types.Config{}
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}
	_, err = conf.Check("test", fset, []*ast.File{file}, info)
	if err != nil {
		t.Fatal(err)
	}

	pkg := &packages.Package{
		Syntax:    []*ast.File{file},
		TypesInfo: info,
	}

	idx := buildFuncIndex([]*packages.Package{pkg})
	if len(idx.byPos) != 2 {
		t.Errorf("expected 2 functions indexed, got %d", len(idx.byPos))
	}
}
