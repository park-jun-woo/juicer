//ff:func feature=scan type=test control=sequence
//ff:what TestBuildGinHField_IdentKeyCov 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestBuildGinHField_IdentKeyCov(t *testing.T) {
	kv := &ast.KeyValueExpr{
		Key:   &ast.Ident{Name: "code"},
		Value: &ast.BasicLit{Kind: token.INT, Value: "200"},
	}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	f := buildGinHField(kv, info)
	if f == nil {
		t.Fatal("expected non-nil")
	}
	if f.Name != "code" {
		t.Fatalf("expected code, got %s", f.Name)
	}
}
