//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestFindAllByType_None 테스트
package nestjs

import "testing"

func TestFindAllByType_None(t *testing.T) {
	src := []byte(`const x = 1;`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	nodes := findAllByType(root, "function_declaration")
	if len(nodes) != 0 {
		t.Fatalf("expected 0, got %d", len(nodes))
	}
}
