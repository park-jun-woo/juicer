//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinRouterType_StarExpr 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinRouterType_StarExpr(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "gin"},
			Sel: &ast.Ident{Name: "Engine"},
		},
	}
	if !isGinRouterType(expr, "gin") {
		t.Fatal("expected true")
	}
}
