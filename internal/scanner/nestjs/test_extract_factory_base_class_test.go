//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractFactoryBaseClass 테스트
package nestjs

import "testing"

func TestExtractFactoryBaseClass(t *testing.T) {
	src := []byte(`const x = PartialType(CreateTaskDto);`)
	root, _ := parseTypeScript(src)
	args := findAllByType(root, "arguments")[0]
	if got := extractFactoryBaseClass(args, src); got != "CreateTaskDto" {
		t.Fatalf("got %q", got)
	}
}
