//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractImports_None 테스트
package fastapi

import "testing"

func TestExtractImports_None(t *testing.T) {
	src := []byte("x = 1\n")
	root, _ := parsePython(src)
	if imports := extractImports(root, src); len(imports) != 0 {
		t.Fatalf("expected none, got %v", imports)
	}
}
