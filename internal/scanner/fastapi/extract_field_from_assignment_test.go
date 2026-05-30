//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractFieldFromAssignment normal/dunder 분기 테스트 (ident==nil은 tree-sitter상 도달 불가)
package fastapi

import "testing"

func TestExtractFieldFromAssignment(t *testing.T) {
	src := []byte("class M:\n    name: str = 'default'\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	assigns := findAllByType(root, "assignment")
	if len(assigns) == 0 {
		t.Fatal("no assignment nodes")
	}
	f := extractFieldFromAssignment(assigns[0], src)
	if f == nil {
		t.Fatal("expected field")
	}
	if f.name != "name" || f.typeName != "str" || !f.hasDefault {
		t.Fatalf("got %+v", f)
	}

	// dunder field
	src2 := []byte("class M:\n    __hidden: int = 0\n")
	root2, _ := parsePython(src2)
	assigns2 := findAllByType(root2, "assignment")
	if len(assigns2) > 0 {
		f2 := extractFieldFromAssignment(assigns2[0], src2)
		if f2 != nil {
			t.Fatal("expected nil for dunder field")
		}
	}
}

func TestExtractFieldFromAssignment_NoType(t *testing.T) {
	// no type annotation -> typeNode absent
	src := []byte("class M:\n    count = 5\n")
	root, _ := parsePython(src)
	assigns := findAllByType(root, "assignment")
	f := extractFieldFromAssignment(assigns[0], src)
	if f == nil || f.name != "count" || f.typeName != "" {
		t.Fatalf("got %+v", f)
	}
}

func TestExtractFieldFromAssignment_FieldCall(t *testing.T) {
	// Field(...) without default kwarg -> hasDefault reset to false
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
