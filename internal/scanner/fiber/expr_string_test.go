//ff:func feature=scan type=test control=selection
//ff:what exprString — AST 표현 문자열화 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"testing"
)

func exprStringFor(t *testing.T, expr string) string {
	t.Helper()
	e, err := parser.ParseExpr(expr)
	if err != nil {
		t.Fatal(err)
	}
	return exprString(e)
}

func TestExprString(t *testing.T) {
	cases := map[string]string{
		"x":              "x",          // Ident
		"pkg.Field":      "pkg.Field",  // SelectorExpr with recv
		"Book{}":         "Book{}",     // CompositeLit
		"*Book":          "*Book",      // StarExpr
		"&req":           "req",        // UnaryExpr
		"make()":         "make()",     // CallExpr
		"m[k]":           "m[k]",       // IndexExpr
		"map[string]int": "map[string]int", // MapType
		"[]byte":         "[]byte",     // ArrayType
		"interface{}":    "interface{}", // InterfaceType
		"42":             "42",         // BasicLit
	}
	for in, want := range cases {
		if got := exprStringFor(t, in); got != want {
			t.Errorf("exprString(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestExprString_Default(t *testing.T) {
	// the default case via an unusual expr type.
	got := exprString(&ast.Ellipsis{})
	if got == "" {
		t.Error("expected non-empty default representation")
	}
}

func TestExprString_CompositeLitNoType(t *testing.T) {
	// inner composite literal in []Book{{...}} has a nil Type -> "{}"
	got := exprStringFor(t, "[]Book{{Title: \"x\"}}")
	if got != "[]Book{}" {
		t.Errorf("outer composite = %q", got)
	}
	// directly test a CompositeLit with nil Type
	cl := &ast.CompositeLit{}
	if s := exprString(cl); s != "{}" {
		t.Errorf("nil-type composite = %q, want {}", s)
	}
}
