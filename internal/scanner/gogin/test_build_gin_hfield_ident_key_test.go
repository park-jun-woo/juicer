//ff:func feature=scan type=test control=sequence
//ff:what TestBuildGinHField_IdentKey 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestBuildGinHField_IdentKey(t *testing.T) {
	kv := &ast.KeyValueExpr{
		Key:   &ast.Ident{Name: "status"},
		Value: &ast.BasicLit{Kind: token.INT, Value: "200"},
	}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	f := buildGinHField(kv, info)
	if f == nil {
		t.Fatal("expected non-nil")
	}
	if f.Name != "status" {
		t.Fatalf("expected 'status', got %q", f.Name)
	}
}
