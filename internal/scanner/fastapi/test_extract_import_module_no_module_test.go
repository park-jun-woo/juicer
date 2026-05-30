//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractImportModule_NoModule 테스트
package fastapi

import "testing"

func TestExtractImportModule_NoModule(t *testing.T) {

	src := []byte("x = 1\n")
	root, _ := parsePython(src)
	if got := extractImportModule(root, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
