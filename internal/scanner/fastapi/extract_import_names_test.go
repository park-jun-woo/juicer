//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractImportNames 테스트
package fastapi

import "testing"

func TestExtractImportNames(t *testing.T) {
	src := []byte("from fastapi import FastAPI, APIRouter\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmts := findAllByType(root, "import_from_statement")
	if len(stmts) == 0 {
		t.Fatal("no import_from_statement")
	}
	names := extractImportNames(stmts[0], src)
	if len(names) < 1 {
		t.Fatalf("expected >= 1 names, got %v", names)
	}
}
