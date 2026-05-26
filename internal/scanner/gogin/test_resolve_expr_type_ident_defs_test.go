//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveExprType_IdentDefs 테스트
package gogin

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveExprType_IdentDefs(t *testing.T) {
	src := `package test

type User struct {
	Name string
}

func f() {
	var u User
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

	// Find 'u' ident where it's USED (not defined)
	var uUseIdent *ast.Ident
	ast.Inspect(file, func(n ast.Node) bool {
		if id, ok := n.(*ast.Ident); ok && id.Name == "u" {
			if _, isDef := info.Defs[id]; !isDef {
				if _, isUse := info.Uses[id]; isUse {
					uUseIdent = id
					return false
				}
			}
		}
		return true
	})

	// Test with Uses[ident] path
	if uUseIdent != nil {
		typeName, _ := resolveExprType(uUseIdent, info)
		if typeName != "User" {
			t.Errorf("expected 'User', got %q", typeName)
		}
	}
}
