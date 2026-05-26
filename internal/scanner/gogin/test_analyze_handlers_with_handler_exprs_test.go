//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeHandlers_WithHandlerExprs 테스트
package gogin

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/packages"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestAnalyzeHandlers_WithHandlerExprs(t *testing.T) {
	src := `package test
func handler() {}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "test.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}

	conf := types.Config{}
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	_, err = conf.Check("test", fset, []*ast.File{file}, info)
	if err != nil {
		t.Fatal(err)
	}

	pkg := &packages.Package{
		Syntax:    []*ast.File{file},
		TypesInfo: info,
	}

	// Create an Ident expr that references the handler function
	handlerIdent := ast.NewIdent("handler")
	endpoints := []scanner.Endpoint{
		{
			Method: "GET",
			Path:   "/test",
		},
	}
	handlerExprsMap := map[int][]ast.Expr{
		0: {handlerIdent},
	}

	analyzeHandlers([]*packages.Package{pkg}, endpoints, ".", handlerExprsMap, nil)
}
