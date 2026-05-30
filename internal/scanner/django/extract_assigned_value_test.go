//ff:func feature=scan type=test control=sequence topic=django
//ff:what extractAssignedValue — 대입문 RHS 값 추출 분기를 검증
package django

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstExprStatement(root *sitter.Node) *sitter.Node {
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == "expression_statement" {
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

func TestExtractAssignedValue_Identifier(t *testing.T) {
	src := []byte("serializer_class = UserSerializer\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := firstExprStatement(root)
	if stmt == nil {
		t.Fatal("no expression_statement")
	}
	got := extractAssignedValue(stmt, "serializer_class", src)
	if got != "UserSerializer" {
		t.Fatalf("got %q, want UserSerializer", got)
	}
}

func TestExtractAssignedValue_Attribute(t *testing.T) {
	src := []byte("serializer_class = mod.UserSerializer\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := firstExprStatement(root)
	got := extractAssignedValue(stmt, "serializer_class", src)
	if got != "mod.UserSerializer" {
		t.Fatalf("got %q, want mod.UserSerializer", got)
	}
}

func TestExtractAssignedValue_WrongName(t *testing.T) {
	src := []byte("other = Value\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := firstExprStatement(root)
	if got := extractAssignedValue(stmt, "serializer_class", src); got != "" {
		t.Fatalf("expected empty for wrong LHS name, got %q", got)
	}
}

func TestExtractAssignedValue_RHSNotIdentOrAttr(t *testing.T) {
	// Correct LHS name but RHS is a literal (no identifier/attribute) -> "".
	src := []byte("serializer_class = 42\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := firstExprStatement(root)
	if got := extractAssignedValue(stmt, "serializer_class", src); got != "" {
		t.Fatalf("expected empty for literal RHS, got %q", got)
	}
}

func TestExtractAssignedValue_NoAssignment(t *testing.T) {
	src := []byte("foo()\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := firstExprStatement(root)
	if got := extractAssignedValue(stmt, "x", src); got != "" {
		t.Fatalf("expected empty for non-assignment, got %q", got)
	}
}
