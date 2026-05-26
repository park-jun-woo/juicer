//ff:func feature=scan type=extract control=sequence
//ff:what TestFindInfoForExpr_NilInfoCanon 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestFindInfoForExpr_NilInfoCanon(t *testing.T) {
	info := &types.Info{}
	_ = info
	result := findInfoForExpr(&ast.Ident{Name: "x"}, nil)
	if result != nil {
		t.Fatal("expected nil")
	}
}
