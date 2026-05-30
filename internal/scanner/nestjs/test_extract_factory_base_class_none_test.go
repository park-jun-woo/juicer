//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractFactoryBaseClass_None 테스트
package nestjs

import "testing"

func TestExtractFactoryBaseClass_None(t *testing.T) {
	src := []byte(`const x = foo('str');`)
	root, _ := parseTypeScript(src)
	args := findAllByType(root, "arguments")[0]
	if got := extractFactoryBaseClass(args, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
