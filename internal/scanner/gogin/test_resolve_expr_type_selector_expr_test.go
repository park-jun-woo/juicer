//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveExprType_SelectorExpr 테스트
package gogin

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveExprType_SelectorExpr(t *testing.T) {
	src := `package test

type Pkg struct{}
type Resp struct {
	Code int
}

var p Pkg
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "test.go", src, 0)
	conf := types.Config{}
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}
	conf.Check("test", fset, []*ast.File{file}, info)

	// Test SelectorExpr path: we need a selector expr in the info.Uses
	selExpr := &ast.SelectorExpr{
		X:   &ast.Ident{Name: "pkg"},
		Sel: &ast.Ident{Name: "Resp"},
	}
	// Without proper Uses mapping, this should return empty
	typeName, fields := resolveExprType(selExpr, info)
	_ = typeName
	_ = fields
}
