//ff:func feature=scan type=extract control=sequence
//ff:what TestFindInfoForExpr_NilTypesCanon 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestFindInfoForExpr_NilTypesCanon(t *testing.T) {
	// findInfoForExpr skips packages with nil TypesInfo
	expr := &ast.Ident{Name: "x"}
	result := findInfoForExpr(expr, nil)
	if result != nil {
		t.Fatal("expected nil")
	}
}
