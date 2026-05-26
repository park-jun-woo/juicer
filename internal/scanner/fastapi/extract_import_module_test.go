//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractImportModule 테스트
package fastapi

import "testing"

func TestExtractImportModule(t *testing.T) {
	src := []byte("from fastapi import FastAPI\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmts := findAllByType(root, "import_from_statement")
	if len(stmts) == 0 {
		t.Fatal("no import_from_statement")
	}
	got := extractImportModule(stmts[0], src)
	if got != "fastapi" {
		t.Fatalf("expected 'fastapi', got %q", got)
	}

	// relative import
	src2 := []byte("from .models import User\n")
	root2, _ := parsePython(src2)
	stmts2 := findAllByType(root2, "import_from_statement")
	got2 := extractImportModule(stmts2[0], src2)
	if got2 == "" {
		t.Fatal("expected non-empty module for relative import")
	}
}
