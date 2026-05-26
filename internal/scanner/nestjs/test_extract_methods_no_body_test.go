//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestExtractMethods_NoBody 테스트
package nestjs

import "testing"

func TestExtractMethods_NoBody(t *testing.T) {
	src := []byte(`declare class Foo;`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	classes := findAllByType(root, "class_declaration")
	for _, cls := range classes {
		result := extractMethods(cls, src, "test.ts")
		if len(result) != 0 {
			t.Fatal("expected 0 for class without body")
		}
	}
}
