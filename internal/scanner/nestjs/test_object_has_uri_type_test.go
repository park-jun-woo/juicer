//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestObjectHasURIType 테스트
package nestjs

import "testing"

func TestObjectHasURIType(t *testing.T) {
	src := []byte(`const o = { type: VersioningType.URI };`)
	root, _ := parseTypeScript(src)
	obj := findAllByType(root, "object")[0]
	if !objectHasURIType(obj, src) {
		t.Fatal("expected true")
	}
	src2 := []byte(`const o = { type: VersioningType.HEADER };`)
	root2, _ := parseTypeScript(src2)
	obj2 := findAllByType(root2, "object")[0]
	if objectHasURIType(obj2, src2) {
		t.Fatal("expected false")
	}
}
