//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractParentClasses_IdentifierAndAttribute 테스트
package django

import "testing"

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

	if parents[1] != "CreateModelMixin" {
		t.Errorf("parents[1] = %q, want CreateModelMixin", parents[1])
	}
}
