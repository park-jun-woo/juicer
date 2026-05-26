//ff:func feature=scan type=test control=sequence
//ff:what TestExprName_Ident 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExprName_Ident(t *testing.T) {
	if exprName(&ast.Ident{Name: "handler"}) != "handler" {
		t.Fatal("ident")
	}
	if exprName(&ast.SelectorExpr{X: &ast.Ident{Name: "h"}, Sel: &ast.Ident{Name: "List"}}) != "h.List" {
		t.Fatal("selector")
	}
	if exprName(&ast.SelectorExpr{X: &ast.CompositeLit{}, Sel: &ast.Ident{Name: "List"}}) != "List" {
		t.Fatal("selector no recv")
	}
	if exprName(&ast.FuncLit{}) != "(inline)" {
		t.Fatal("funclit")
	}
	if exprName(&ast.CallExpr{Fun: &ast.Ident{Name: "f"}}) != "f()" {
		t.Fatal("callexpr")
	}
	if exprName(&ast.BasicLit{}) != "" {
		t.Fatal("unknown")
	}
}

