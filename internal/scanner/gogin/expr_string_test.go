//ff:func feature=scan type=test control=sequence
//ff:what TestExprString_Ident 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExprString_Ident(t *testing.T) {
	if exprString(&ast.Ident{Name: "x"}) != "x" {
		t.Fatal("ident")
	}
	if exprString(&ast.SelectorExpr{X: &ast.Ident{Name: "h"}, Sel: &ast.Ident{Name: "F"}}) != "h.F" {
		t.Fatal("selector")
	}
	if exprString(&ast.CompositeLit{Type: &ast.Ident{Name: "T"}}) != "T{}" {
		t.Fatal("composite")
	}
	if exprString(&ast.CompositeLit{}) != "{}" {
		t.Fatal("composite no type")
	}
	if exprString(&ast.StarExpr{X: &ast.Ident{Name: "T"}}) != "*T" {
		t.Fatal("star")
	}
	if exprString(&ast.UnaryExpr{X: &ast.Ident{Name: "x"}}) != "x" {
		t.Fatal("unary")
	}
	if exprString(&ast.CallExpr{Fun: &ast.Ident{Name: "f"}}) != "f()" {
		t.Fatal("call")
	}
	if exprString(&ast.IndexExpr{X: &ast.Ident{Name: "m"}, Index: &ast.Ident{Name: "k"}}) != "m[k]" {
		t.Fatal("index")
	}
	if exprString(&ast.MapType{Key: &ast.Ident{Name: "string"}, Value: &ast.Ident{Name: "int"}}) != "map[string]int" {
		t.Fatal("map")
	}
	if exprString(&ast.ArrayType{Elt: &ast.Ident{Name: "int"}}) != "[]int" {
		t.Fatal("array")
	}
	if exprString(&ast.InterfaceType{}) != "interface{}" {
		t.Fatal("interface")
	}
	if exprString(&ast.BasicLit{Value: "42"}) != "42" {
		t.Fatal("basic")
	}
}

