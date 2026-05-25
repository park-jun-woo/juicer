//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestResolveCallerArgs_FewerArgs 테스트
package scanner

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveCallerArgs_FewerArgs(t *testing.T) {
	src := `package test

func helper(a int, b int) {}

func caller() {
	helper(200)
}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "test.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}

	conf := types.Config{
		Error: func(err error) {}, // ignore errors for incomplete args
	}
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	conf.Check("test", fset, []*ast.File{file}, info)

	var helperDecl *ast.FuncDecl
	var callerDecl *ast.FuncDecl
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if fn.Name.Name == "helper" {
			helperDecl = fn
		}
		if fn.Name.Name == "caller" {
			callerDecl = fn
		}
	}

	if helperDecl == nil || callerDecl == nil {
		t.Skip("could not parse functions due to type check errors")
	}

	var callExpr *ast.CallExpr
	ast.Inspect(callerDecl.Body, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			callExpr = c
			return false
		}
		return true
	})

	if callExpr == nil {
		t.Skip("could not find call expression")
	}

	status, _, _, _ := resolveCallerArgs(helperDecl, callExpr, info, info)
	// Should handle fewer args than params gracefully
	_ = status
}
