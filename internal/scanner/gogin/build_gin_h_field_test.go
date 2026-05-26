//ff:func feature=scan type=test control=sequence
//ff:what TestBuildGinHField_NonKV 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestBuildGinHField_NonKV(t *testing.T) {
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	f := buildGinHField(&ast.Ident{Name: "x"}, info)
	if f != nil {
		t.Fatal("expected nil")
	}

	// KV with string key
	kv := &ast.KeyValueExpr{
		Key:   &ast.BasicLit{Kind: token.STRING, Value: `"name"`},
		Value: &ast.BasicLit{Kind: token.STRING, Value: `"hello"`},
	}
	f = buildGinHField(kv, info)
	if f == nil {
		t.Fatal("expected non-nil")
	}
	if f.JSON != "name" {
		t.Fatalf("expected name, got %s", f.JSON)
	}

	// KV with ident key
	kv2 := &ast.KeyValueExpr{
		Key:   &ast.Ident{Name: "status"},
		Value: &ast.BasicLit{Kind: token.INT, Value: "200"},
	}
	f2 := buildGinHField(kv2, info)
	if f2 == nil || f2.JSON != "status" {
		t.Fatal("expected status field")
	}

	// KV with empty key (non-string, non-ident)
	kv3 := &ast.KeyValueExpr{
		Key:   &ast.CompositeLit{},
		Value: &ast.BasicLit{Kind: token.STRING, Value: `"x"`},
	}
	f3 := buildGinHField(kv3, info)
	if f3 != nil {
		t.Fatal("expected nil for empty key")
	}

	// KV with nested gin.H value
	nestedGinH := &ast.CompositeLit{
		Type: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "gin"},
			Sel: &ast.Ident{Name: "H"},
		},
		Elts: []ast.Expr{
			&ast.KeyValueExpr{
				Key:   &ast.BasicLit{Kind: token.STRING, Value: `"nested"`},
				Value: &ast.BasicLit{Kind: token.STRING, Value: `"value"`},
			},
		},
	}
	kv4 := &ast.KeyValueExpr{
		Key:   &ast.BasicLit{Kind: token.STRING, Value: `"data"`},
		Value: nestedGinH,
	}
	f4 := buildGinHField(kv4, info)
	if f4 == nil {
		t.Fatal("expected non-nil for nested gin.H")
	}
	if f4.Type != "object" {
		t.Fatalf("expected object type, got %s", f4.Type)
	}
}
