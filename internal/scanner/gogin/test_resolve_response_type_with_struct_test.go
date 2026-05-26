//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what TestResolveResponseType_WithStruct 테스트
package gogin

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveResponseType_WithStruct(t *testing.T) {
	src := `package test

type Resp struct {
	Code int
	Msg  string
}

var r Resp
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

	var rIdent *ast.Ident
	for ident := range info.Defs {
		if ident.Name == "r" {
			rIdent = ident
			break
		}
	}

	typeName, fields, confidence := resolveResponseType(rIdent, info)
	if typeName != "Resp" {
		t.Errorf("expected 'Resp', got %q", typeName)
	}
	if confidence != "full" {
		t.Errorf("expected 'full', got %q", confidence)
	}
	if len(fields) != 2 {
		t.Errorf("expected 2 fields, got %d", len(fields))
	}
}
