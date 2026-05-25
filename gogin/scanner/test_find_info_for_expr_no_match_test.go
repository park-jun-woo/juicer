//ff:func feature=scan type=extract control=sequence
//ff:what TestFindInfoForExpr_NoMatch 테스트
package scanner

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestFindInfoForExpr_NoMatch(t *testing.T) {
	// Expr with position outside any file
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

	// Create an expr with a position that doesn't match any file
	outsideIdent := &ast.Ident{NamePos: 99999, Name: "x"}
	result := findInfoForExpr(outsideIdent, []*packages.Package{pkg})
	if result != nil {
		t.Error("expected nil for expr outside all files")
	}
}
