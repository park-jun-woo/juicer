//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractFactoryStringArray_None 테스트
package nestjs

import "testing"

func TestExtractFactoryStringArray_None(t *testing.T) {
	src := []byte(`const x = PartialType(Base);`)
	root, _ := parseTypeScript(src)
	args := findAllByType(root, "arguments")[0]
	if got := extractFactoryStringArray(args, src); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}
