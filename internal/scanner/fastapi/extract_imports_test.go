//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractImports 테스트
package fastapi

import "testing"

func TestExtractImports(t *testing.T) {
	src := []byte("from .models import User, Order\nfrom fastapi import FastAPI\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	imports := extractImports(root, src)
	if len(imports) < 1 {
		t.Fatalf("expected >= 1 imports, got %v", imports)
	}
}
