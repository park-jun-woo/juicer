//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestResolveFieldsWithInheritance 테스트
package fastapi

import "testing"

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

	fields := resolveFieldsWithInheritance(child, root, src, map[string]bool{})
	names := map[string]bool{}
	for _, f := range fields {
		names[f.name] = true
	}
	if !names["id"] || !names["name"] {
		t.Fatalf("expected id+name merged, got %v", names)
	}
}
