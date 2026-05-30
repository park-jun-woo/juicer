//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractFactoryStringArray 테스트
package nestjs

import "testing"

func TestExtractFactoryStringArray(t *testing.T) {
	src := []byte(`const x = OmitType(Base, ['a', 'b']);`)
	root, _ := parseTypeScript(src)
	args := findAllByType(root, "arguments")[0]
	got := extractFactoryStringArray(args, src)
	if len(got) != 2 || got[0] != "a" {
		t.Fatalf("got %v", got)
	}
}
