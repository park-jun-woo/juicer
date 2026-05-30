//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what resolveFieldsWithInheritance 테스트
package fastapi

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func classByName(root *sitter.Node, src []byte, name string) *sitter.Node {
	for _, c := range findAllByType(root, "class_definition") {
		id := findChildByType(c, "identifier")
		if id != nil && nodeText(id, src) == name {
			return c
		}
	}
	return nil
}

func TestResolveFieldsWithInheritance(t *testing.T) {
	src := []byte(`class Base(BaseModel):
    id: int

class Child(Base):
    name: str
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	child := classByName(root, src, "Child")
	if child == nil {
		t.Fatal("Child class not found")
	}
	// Base is a well-known base (BaseModel skipped), and Base merged in.
	fields := resolveFieldsWithInheritance(child, root, src, map[string]bool{})
	names := map[string]bool{}
	for _, f := range fields {
		names[f.name] = true
	}
	if !names["id"] || !names["name"] {
		t.Fatalf("expected id+name merged, got %v", names)
	}
}

func TestResolveFieldsWithInheritance_Cycle(t *testing.T) {
	src := []byte("class A(B):\n    x: int\n")
	root, _ := parsePython(src)
	a := classByName(root, src, "A")
	// pre-mark A as visited -> immediate nil
	if got := resolveFieldsWithInheritance(a, root, src, map[string]bool{"A": true}); got != nil {
		t.Fatalf("expected nil for visited cycle, got %v", got)
	}
}

func TestResolveFieldsWithInheritance_WellKnownParentSkipped(t *testing.T) {
	src := []byte("class M(BaseModel):\n    a: int\n")
	root, _ := parsePython(src)
	m := classByName(root, src, "M")
	fields := resolveFieldsWithInheritance(m, root, src, map[string]bool{})
	if len(fields) != 1 || fields[0].name != "a" {
		t.Fatalf("expected only own field a, got %v", fields)
	}
}
