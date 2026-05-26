//ff:func feature=scan type=extract control=sequence
//ff:what TestFindInfoForExpr 테스트
package gogin

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestFindInfoForExpr(t *testing.T) {
	src := `package test
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

	// Find the Hello function's name ident
	fn := file.Decls[0].(*ast.FuncDecl)
	result := findInfoForExpr(fn.Name, []*packages.Package{pkg})
	if result == nil {
		t.Error("expected to find info for expr")
	}
}
