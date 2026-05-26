//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveExprType_CompositeLit 테스트
package gogin

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveExprType_CompositeLit(t *testing.T) {
	src := `package test

type User struct {
	Name string
}

func f() {
	u := User{Name: "test"}
	_ = u
}
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

	// Find the composite literal
	var compLit *ast.CompositeLit
	ast.Inspect(file, func(n ast.Node) bool {
		if cl, ok := n.(*ast.CompositeLit); ok {
			compLit = cl
			return false
		}
		return true
	})

	if compLit != nil {
		typeName, fields := resolveExprType(compLit, info)
		if typeName != "User" {
			t.Errorf("expected 'User', got %q", typeName)
		}
		if len(fields) != 1 {
			t.Errorf("expected 1 field, got %d", len(fields))
		}
	}
}
