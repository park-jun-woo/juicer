//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractEnumArray 테스트
package nestjs

import "testing"

func TestExtractEnumArray(t *testing.T) {
	src := []byte(`const o = { enum: ['a', 'b'] };`)
	root, _ := parseTypeScript(src)
	obj := findAllByType(root, "object")[0]
	vals := extractEnumArray(obj, src)
	if len(vals) != 2 {
		t.Fatalf("got %v", vals)
	}
}
