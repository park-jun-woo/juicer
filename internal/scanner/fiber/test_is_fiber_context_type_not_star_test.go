//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberContextType_NotStar 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberContextType_NotStar(t *testing.T) {
	expr := &ast.SelectorExpr{
		X:   &ast.Ident{Name: "fiber"},
		Sel: &ast.Ident{Name: "Ctx"},
	}
	if isFiberContextType(expr) {
		t.Fatal("expected false — must be *fiber.Ctx with star")
	}
}

func TestIsFiberContextType_Match(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "fiber"},
			Sel: &ast.Ident{Name: "Ctx"},
		},
	}
	if !isFiberContextType(expr) {
		t.Fatal("expected true for *fiber.Ctx")
	}
}

func TestIsFiberContextType_StarNotSelector(t *testing.T) {
	expr := &ast.StarExpr{X: &ast.Ident{Name: "Ctx"}}
	if isFiberContextType(expr) {
		t.Fatal("expected false — *Ctx without selector")
	}
}

func TestIsFiberContextType_SelectorNotIdent(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.CallExpr{Fun: &ast.Ident{Name: "f"}},
			Sel: &ast.Ident{Name: "Ctx"},
		},
	}
	if isFiberContextType(expr) {
		t.Fatal("expected false — selector X not an ident")
	}
}

func TestIsFiberContextType_WrongNames(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "http"},
			Sel: &ast.Ident{Name: "Request"},
		},
	}
	if isFiberContextType(expr) {
		t.Fatal("expected false for *http.Request")
	}
}
