//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what tryImportedName 테스트
package fastapi

import "testing"

func TestTryImportedName(t *testing.T) {
	src := []byte("from fastapi import FastAPI, APIRouter\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmts := findAllByType(root, "import_from_statement")
	if len(stmts) == 0 {
		t.Fatal("no import_from_statement")
	}
	stmt := stmts[0]

	// First child (index 0) should return ""
	got := tryImportedName(stmt.Child(0), stmt, 0, src)
	if got != "" {
		t.Fatalf("expected empty for index 0, got %q", got)
	}

	// Find an identifier child after "import"
	found := false
	for i := 1; i < int(stmt.ChildCount()); i++ {
		child := stmt.Child(i)
		name := tryImportedName(child, stmt, i, src)
		if name == "FastAPI" || name == "APIRouter" {
			found = true
		}
	}
	if !found {
		t.Fatal("did not find any imported name")
	}
}
