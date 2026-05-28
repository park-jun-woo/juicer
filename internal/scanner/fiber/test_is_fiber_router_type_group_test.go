//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberRouterType_Group 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberRouterType_Group(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "fiber"},
			Sel: &ast.Ident{Name: "Group"},
		},
	}
	if !isFiberRouterType(expr, "fiber") {
		t.Fatal("expected true")
	}
}
