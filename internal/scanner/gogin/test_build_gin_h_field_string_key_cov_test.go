//ff:func feature=scan type=test control=sequence
//ff:what TestBuildGinHField_StringKeyCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestBuildGinHField_StringKeyCov(t *testing.T) {
	kv := &ast.KeyValueExpr{
		Key:   &ast.BasicLit{Kind: token.STRING, Value: `"name"`},
		Value: &ast.BasicLit{Kind: token.STRING, Value: `"hello"`},
	}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	f := buildGinHField(kv, info)
	if f == nil {
		t.Fatal("expected non-nil")
	}
	if f.Name != "name" {
		t.Fatalf("expected name, got %s", f.Name)
	}
}
