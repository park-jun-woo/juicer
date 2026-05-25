package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveBindType_EmptyArgs(t *testing.T) {
	call := &ast.CallExpr{}
	info := &types.Info{}
	name, fields := resolveBindType(call, info)
	if name != "" || fields != nil {
		t.Fatal("expected empty")
	}
}

func TestResolveBindType_NilInfoCase(t *testing.T) {
	call := &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "req"}}}
	name, fields := resolveBindType(call, nil)
	if name != "" || fields != nil {
		t.Fatal("expected empty with nil info")
	}
}
