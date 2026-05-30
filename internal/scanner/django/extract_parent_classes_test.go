//ff:func feature=scan type=test control=sequence topic=django
//ff:what extractParentClasses — 부모 클래스명 추출 분기를 검증
package django

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstClassDef(root *sitter.Node) *sitter.Node {
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == "class_definition" {
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

func TestExtractParentClasses_IdentifierAndAttribute(t *testing.T) {
	src := []byte("class V(ModelViewSet, mixins.CreateModelMixin):\n    pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	cd := firstClassDef(root)
	if cd == nil {
		t.Fatal("no class_definition")
	}
	parents := extractParentClasses(cd, src)
	if len(parents) != 2 {
		t.Fatalf("expected 2 parents, got %v", parents)
	}
	if parents[0] != "ModelViewSet" {
		t.Errorf("parents[0] = %q, want ModelViewSet", parents[0])
	}
	// attribute "mixins.CreateModelMixin" -> last segment.
	if parents[1] != "CreateModelMixin" {
		t.Errorf("parents[1] = %q, want CreateModelMixin", parents[1])
	}
}

func TestExtractParentClasses_NoParents(t *testing.T) {
	src := []byte("class C:\n    pass\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	cd := firstClassDef(root)
	if p := extractParentClasses(cd, src); p != nil {
		t.Fatalf("expected nil for class without bases, got %v", p)
	}
}
