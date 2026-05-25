//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestResolveExprType_WithTypedCode 테스트
package scanner

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveExprType_WithTypedCode(t *testing.T) {
	src := `package test

type User struct {
	Name string
	Age  int
}

var u User
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

	// Find the 'u' ident in the var declaration
	var uIdent *ast.Ident
	for ident := range info.Defs {
		if ident.Name == "u" {
			uIdent = ident
			break
		}
	}
	if uIdent == nil {
		t.Fatal("could not find 'u' ident")
	}

	typeName, fields := resolveExprType(uIdent, info)
	if typeName != "User" {
		t.Errorf("expected typeName 'User', got %q", typeName)
	}
	if len(fields) != 2 {
		t.Errorf("expected 2 fields, got %d", len(fields))
	}
}
