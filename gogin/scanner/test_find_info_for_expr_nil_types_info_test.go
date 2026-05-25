//ff:func feature=scan type=extract control=sequence
//ff:what TestFindInfoForExpr_NilTypesInfo 테스트
package scanner

import (
	"go/ast"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestFindInfoForExpr_NilTypesInfo(t *testing.T) {
	pkg := &packages.Package{
		TypesInfo: nil,
	}
	ident := &ast.Ident{NamePos: 1, Name: "x"}
	result := findInfoForExpr(ident, []*packages.Package{pkg})
	if result != nil {
		t.Error("expected nil for nil TypesInfo")
	}
}
