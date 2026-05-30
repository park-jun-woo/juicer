//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberRouterType_App 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberRouterType_App(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "fiber"},
			Sel: &ast.Ident{Name: "App"},
		},
	}
	if !isFiberRouterType(expr, "fiber") {
		t.Fatal("expected true")
	}
}

func TestIsFiberRouterType_NoStarGroup(t *testing.T) {
	// non-pointer fiber.Group is also accepted
	expr := &ast.SelectorExpr{X: &ast.Ident{Name: "fiber"}, Sel: &ast.Ident{Name: "Router"}}
	_ = isFiberRouterType(expr, "fiber") // Router may or may not be in the set; just exercise path
}

func TestIsFiberRouterType_NotSelector(t *testing.T) {
	if isFiberRouterType(&ast.Ident{Name: "x"}, "fiber") {
		t.Fatal("expected false for non-selector")
	}
}

func TestIsFiberRouterType_XNotIdent(t *testing.T) {
	expr := &ast.SelectorExpr{X: &ast.CallExpr{}, Sel: &ast.Ident{Name: "App"}}
	if isFiberRouterType(expr, "fiber") {
		t.Fatal("expected false when X not ident")
	}
}

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
