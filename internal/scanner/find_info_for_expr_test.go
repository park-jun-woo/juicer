//ff:func feature=scan type=test control=sequence
//ff:what TestFindInfoForExpr_NilPkgs 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestFindInfoForExpr_NilPkgs(t *testing.T) {
	expr := &ast.Ident{Name: "x"}
	result := findInfoForExpr(expr, nil)
	if result != nil {
		t.Fatal("expected nil")
	}
}
