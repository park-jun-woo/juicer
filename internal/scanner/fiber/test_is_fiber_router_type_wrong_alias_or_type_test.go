//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberRouterType_WrongAliasOrType 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberRouterType_WrongAliasOrType(t *testing.T) {
	expr := &ast.SelectorExpr{X: &ast.Ident{Name: "other"}, Sel: &ast.Ident{Name: "App"}}
	if isFiberRouterType(expr, "fiber") {
		t.Fatal("expected false for wrong alias")
	}
	expr2 := &ast.SelectorExpr{X: &ast.Ident{Name: "fiber"}, Sel: &ast.Ident{Name: "Ctx"}}
	if isFiberRouterType(expr2, "fiber") {
		t.Fatal("expected false for non-router type")
	}
}
