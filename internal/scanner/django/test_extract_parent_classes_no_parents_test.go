//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractParentClasses_NoParents 테스트
package django

import "testing"

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
