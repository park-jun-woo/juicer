package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveExprType_NilInfoCase(t *testing.T) {
	tn, fields := resolveExprType(&ast.Ident{Name: "x"}, nil)
	if tn != "" || fields != nil {
		t.Fatal("expected empty")
	}
}

func TestResolveExprType_IdentNotFound(t *testing.T) {
	info := &types.Info{
		Uses: make(map[*ast.Ident]types.Object),
		Defs: make(map[*ast.Ident]types.Object),
	}
	tn, fields := resolveExprType(&ast.Ident{Name: "x"}, info)
	if tn != "" || fields != nil {
		t.Fatal("expected empty")
	}
}
