//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what collectParentNames: 부모 클래스 추출 / keyword arg 스킵 / argument_list 없음
package fastapi

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstClass(t *testing.T, src []byte) (*sitter.Node, []byte) {
	t.Helper()
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	cls := findAllByType(root, "class_definition")
	if len(cls) == 0 {
		t.Fatal("no class_definition")
	}
	return cls[0], src
}

func TestCollectParentNames_WithParents(t *testing.T) {
	cls, src := firstClass(t, []byte("class User(Base, BaseModel, table=True): pass\n"))
	names := collectParentNames(cls, src)
	if len(names) != 2 || names[0] != "Base" || names[1] != "BaseModel" {
		t.Fatalf("got %v", names)
	}
}

func TestCollectParentNames_NoArgs(t *testing.T) {
	cls, src := firstClass(t, []byte("class Plain: pass\n"))
	if names := collectParentNames(cls, src); names != nil {
		t.Fatalf("expected nil, got %v", names)
	}
}
