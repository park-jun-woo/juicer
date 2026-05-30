//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractFieldFromAssignment_FieldCall 테스트
package fastapi

import "testing"

func TestExtractFieldFromAssignment_FieldCall(t *testing.T) {

	src := []byte("class M:\n    age: int = Field(ge=0)\n")
	root, _ := parsePython(src)
	assigns := findAllByType(root, "assignment")
	f := extractFieldFromAssignment(assigns[0], src)
	if f == nil || f.hasDefault {
		t.Fatalf("expected hasDefault=false, got %+v", f)
	}
	if f.ge == nil || *f.ge != 0 {
		t.Fatalf("ge wrong: %+v", f)
	}
}
