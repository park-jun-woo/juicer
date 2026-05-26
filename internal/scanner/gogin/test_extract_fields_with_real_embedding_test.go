//ff:func feature=scan type=extract control=iteration dimension=2
//ff:what TestExtractFields_WithRealEmbedding 테스트
package gogin

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestExtractFields_WithRealEmbedding(t *testing.T) {
	src := `package test

type Base struct {
	ID int
}

type User struct {
	Base
	Name string
}

var u User
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

	var uIdent *ast.Ident
	for ident := range info.Defs {
		if ident.Name == "u" {
			uIdent = ident
			break
		}
	}

	if uIdent != nil {
		typeName, fields := resolveExprType(uIdent, info)
		if typeName != "User" {
			t.Errorf("expected 'User', got %q", typeName)
		}
		// Should have ID (from Base) and Name
		if len(fields) != 2 {
			t.Errorf("expected 2 fields (ID from Base + Name), got %d", len(fields))
			for _, f := range fields {
				t.Logf("  field: %s %s", f.Name, f.Type)
			}
		}
	}
}
