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

func TestExtractImports_None(t *testing.T) {
	src := []byte("x = 1\n")
	root, _ := parsePython(src)
	if imports := extractImports(root, src); len(imports) != 0 {
		t.Fatalf("expected none, got %v", imports)
	}
}

func TestExtractImports_MultipleNamesAndModules(t *testing.T) {
	src := []byte("from app.models import User, Item\nfrom .routes import router\n")
	root, _ := parsePython(src)
	imports := extractImports(root, src)
	if len(imports) != 3 {
		t.Fatalf("expected 3, got %d: %+v", len(imports), imports)
	}
}
