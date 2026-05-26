//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what identText 테스트
package fastapi

import "testing"

func TestIdentText(t *testing.T) {
	src := []byte("x = 1\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := findChildByType(root, "expression_statement")
	if stmt == nil {
		t.Fatal("no expression_statement")
	}
	assign := stmt.Child(0)
	name := identText(assign, src)
	if name != "x" {
		t.Fatalf("expected 'x', got %q", name)
	}

	// node without identifier
	got := identText(root, src)
	if got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
