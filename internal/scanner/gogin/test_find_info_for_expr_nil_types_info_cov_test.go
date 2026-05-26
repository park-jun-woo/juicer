//ff:func feature=scan type=test control=sequence
//ff:what TestFindInfoForExpr_NilTypesInfoCov 테스트
package gogin

import (
	"go/ast"
	"testing"
	"golang.org/x/tools/go/packages"
)

func TestFindInfoForExpr_NilTypesInfoCov(t *testing.T) {
	expr := &ast.Ident{Name: "x"}
	pkg := &packages.Package{}
	result := findInfoForExpr(expr, []*packages.Package{pkg})
	if result != nil {
		t.Fatal("expected nil")
	}
}
