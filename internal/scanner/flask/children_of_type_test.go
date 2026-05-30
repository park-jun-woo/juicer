//ff:func feature=scan type=test control=sequence topic=flask
//ff:what childrenOfType 테스트
package flask

import "testing"

func TestChildrenOfType(t *testing.T) {
	root, err := parsePython([]byte("a = 1\nb = 2\n"))
	if err != nil {
		t.Fatal(err)
	}
	// module -> expression_statement children
	stmts := childrenOfType(root, "expression_statement")
	if len(stmts) != 2 {
		t.Fatalf("expected 2 expression_statements, got %d", len(stmts))
	}
	if got := childrenOfType(root, "function_definition"); len(got) != 0 {
		t.Fatalf("expected 0 functions, got %d", len(got))
	}
}
