//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestResolveCallerArgs_WithTypedParams 테스트
package gogin

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveCallerArgs_WithTypedParams(t *testing.T) {
	// Build a simple Go source with type-checked functions
	src := `package test

func helper(code int, data interface{}) {}

func caller() {
	helper(200, "hello")
}
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

	// Find the helper FuncDecl and the call in caller
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
		t.Fatal("expected to find both functions")
	}

	// Find the call expression in caller body
	var callExpr *ast.CallExpr
	ast.Inspect(callerDecl.Body, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			callExpr = c
			return false
		}
		return true
	})

	if callExpr == nil {
		t.Fatal("expected to find call expression")
	}

	status, _, _, _ := resolveCallerArgs(helperDecl, callExpr, info, info)
	if status != "200" {
		t.Errorf("expected status '200', got %q", status)
	}
}
