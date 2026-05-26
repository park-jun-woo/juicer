//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinRouterType_NotIdentXCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinRouterType_NotIdentXCov(t *testing.T) {
	sel := &ast.SelectorExpr{X: &ast.CompositeLit{}, Sel: &ast.Ident{Name: "Engine"}}
	if isGinRouterType(sel, "gin") {
		t.Fatal("expected false")
	}
}
