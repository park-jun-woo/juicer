//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeHandlers_WithHandlerExprs 테스트
package scanner

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/packages"
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
	// Set its position within the file range
	endpoints := []Endpoint{
		{
			Method:       "GET",
			Path:         "/test",
			handlerExprs: []ast.Expr{handlerIdent},
		},
	}

	analyzeHandlers([]*packages.Package{pkg}, endpoints, ".")
	// After analysis, handlerExprs should be nil (cleaned up)
	if endpoints[0].handlerExprs != nil {
		t.Error("expected handlerExprs to be nil after analysis")
	}
}
