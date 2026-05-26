//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestResolveResponseType_NameOnly 테스트
package gogin

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveResponseType_NameOnly(t *testing.T) {
	// A type that resolves to a typeName but no fields (e.g., basic named type)
	src := `package test

type Status int

var s Status
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

	var sIdent *ast.Ident
	for ident := range info.Defs {
		if ident.Name == "s" {
			sIdent = ident
			break
		}
	}

	if sIdent != nil {
		typeName, fields, confidence := resolveResponseType(sIdent, info)
		if typeName != "Status" {
			t.Errorf("expected 'Status', got %q", typeName)
		}
		if fields != nil {
			t.Error("expected nil fields for non-struct type")
		}
		if confidence != "" {
			t.Errorf("expected empty confidence, got %q", confidence)
		}
	}
}
