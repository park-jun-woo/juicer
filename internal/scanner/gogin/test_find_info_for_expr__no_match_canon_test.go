//ff:func feature=scan type=extract control=sequence
//ff:what TestFindInfoForExpr_NoMatchCanon 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestFindInfoForExpr_NoMatchCanon(t *testing.T) {
	result := findInfoForExpr(&ast.Ident{Name: "x"}, nil)
	if result != nil {
		t.Fatal("expected nil")
	}
}
