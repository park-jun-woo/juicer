//ff:func feature=scan type=extract control=sequence
//ff:what TestIsGinRouterType_NotGin 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinRouterType_NotGin(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "http"},
			Sel: &ast.Ident{Name: "Server"},
		},
	}
	if isGinRouterType(expr, "gin") {
		t.Fatal("expected false")
	}
}
