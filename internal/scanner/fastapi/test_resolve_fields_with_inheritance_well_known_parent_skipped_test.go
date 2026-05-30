//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestResolveFieldsWithInheritance_WellKnownParentSkipped 테스트
package fastapi

import "testing"

func TestResolveFieldsWithInheritance_WellKnownParentSkipped(t *testing.T) {
	src := []byte("class M(BaseModel):\n    a: int\n")
	root, _ := parsePython(src)
	m := classByName(root, src, "M")
	fields := resolveFieldsWithInheritance(m, root, src, map[string]bool{})
	if len(fields) != 1 || fields[0].name != "a" {
		t.Fatalf("expected only own field a, got %v", fields)
	}
}
