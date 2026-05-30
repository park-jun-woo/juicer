//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractFieldCallConstraints: Field(...) 제약 추출 / 비Field / call없음 분기
package fastapi

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstAssignment(t *testing.T, src []byte) (*sitter.Node, []byte) {
	t.Helper()
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	a := findAllByType(root, "assignment")
	if len(a) == 0 {
		t.Fatal("no assignment")
	}
	return a[0], src
}

func TestExtractFieldCallConstraints_Field(t *testing.T) {
	assign, src := firstAssignment(t, []byte("age: int = Field(ge=0, le=120)\n"))
	f := &pydanticField{}
	extractFieldCallConstraints(assign, src, f)
	if f.ge == nil || *f.ge != 0 || f.le == nil || *f.le != 120 {
		t.Fatalf("got %+v", f)
	}
}

func TestExtractFieldCallConstraints_NotField(t *testing.T) {
	assign, src := firstAssignment(t, []byte("x: int = other(ge=0)\n"))
	f := &pydanticField{}
	extractFieldCallConstraints(assign, src, f)
	if f.ge != nil {
		t.Fatalf("expected no constraints, got %+v", f)
	}
}

func TestExtractFieldCallConstraints_NoCall(t *testing.T) {
	assign, src := firstAssignment(t, []byte("x: int = 5\n"))
	f := &pydanticField{}
	extractFieldCallConstraints(assign, src, f)
	if f.ge != nil || f.hasDefault {
		t.Fatalf("expected no-op, got %+v", f)
	}
	// the "call without argument_list" branch is unreachable: a tree-sitter
	// `call` node always has an argument_list child.
}
