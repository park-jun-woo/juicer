//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what childrenOfType 테스트
package fastapi

import "testing"

func TestChildrenOfType(t *testing.T) {
	src := []byte("x = 1\ny = 2\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmts := childrenOfType(root, "expression_statement")
	if len(stmts) != 2 {
		t.Fatalf("expected 2 expression_statements, got %d", len(stmts))
	}
	none := childrenOfType(root, "nonexistent_type")
	if len(none) != 0 {
		t.Fatalf("expected 0, got %d", len(none))
	}
}
