//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractEnumArray_NoEnumKey 테스트
package nestjs

import "testing"

func TestExtractEnumArray_NoEnumKey(t *testing.T) {
	src := []byte(`const o = { other: ['a'] };`)
	root, _ := parseTypeScript(src)
	obj := findAllByType(root, "object")[0]
	if vals := extractEnumArray(obj, src); vals != nil {
		t.Fatalf("expected nil, got %v", vals)
	}
}
