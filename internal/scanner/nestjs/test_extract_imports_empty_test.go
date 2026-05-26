//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractImports_Empty 테스트
package nestjs

import "testing"

func TestExtractImports_Empty(t *testing.T) {
	src := []byte(`const x = 1;`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	result := extractImports(root, src)
	if len(result) != 0 {
		t.Fatalf("expected empty, got %v", result)
	}
}
