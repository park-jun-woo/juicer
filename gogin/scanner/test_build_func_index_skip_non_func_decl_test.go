//ff:func feature=scan type=extract control=sequence
//ff:what TestBuildFuncIndex_SkipNonFuncDecl 테스트
package scanner

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestBuildFuncIndex_SkipNonFuncDecl(t *testing.T) {
	src := `package test
var x = 1
func Hello() {}
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
	if len(idx.byPos) != 1 {
		t.Errorf("expected 1 function indexed, got %d", len(idx.byPos))
	}
}
