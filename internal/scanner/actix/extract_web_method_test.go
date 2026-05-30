//ff:func feature=scan type=test control=sequence topic=actix
//ff:what extractWebMethod — field_expression에서 HTTP 메서드 추출 분기를 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstFieldExpr(root *sitter.Node) *sitter.Node {
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == "field_expression" {
			found = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	return found
}

func TestExtractWebMethod(t *testing.T) {
	// web::get().to(h): the field_expression "web::get().to" contains the
	// web::get() call_expression child.
	src := []byte(`fn f() { web::get().to(h); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fe := firstFieldExpr(root)
	if fe == nil {
		t.Fatal("no field_expression found")
	}
	if got := extractWebMethod(fe, src); got != "GET" {
		t.Fatalf("extractWebMethod = %q, want GET", got)
	}
}

func TestExtractWebMethod_None(t *testing.T) {
	// foo.bar field_expression has no web::<method> call child -> "".
	src := []byte(`fn f() { foo.bar; }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fe := firstFieldExpr(root)
	if fe == nil {
		t.Fatal("no field_expression found")
	}
	if got := extractWebMethod(fe, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
