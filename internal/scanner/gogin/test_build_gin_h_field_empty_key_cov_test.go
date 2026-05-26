//ff:func feature=scan type=test control=sequence
//ff:what TestBuildGinHField_EmptyKeyCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestBuildGinHField_EmptyKeyCov(t *testing.T) {
	kv := &ast.KeyValueExpr{
		Key:   &ast.CompositeLit{},
		Value: &ast.BasicLit{Kind: token.STRING, Value: `"x"`},
	}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	f := buildGinHField(kv, info)
	if f != nil {
		t.Fatal("expected nil for empty key")
	}
}
