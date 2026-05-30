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

func TestTryImportedName_DottedName(t *testing.T) {
	// "import os.path" -> dotted_name child preceded by "import"
	src := []byte("import os.path\n")
	root, _ := parsePython(src)
	stmts := findAllByType(root, "import_statement")
	if len(stmts) == 0 {
		t.Skip("no import_statement")
	}
	stmt := stmts[0]
	got := ""
	for i := 1; i < int(stmt.ChildCount()); i++ {
		if name := tryImportedName(stmt.Child(i), stmt, i, src); name != "" {
			got = name
		}
	}
	if got != "os.path" {
		t.Fatalf("dotted_name import: got %q", got)
	}
}

func TestTryImportedName_NotPrecededByImport(t *testing.T) {
	// "from fastapi import X": the module "fastapi" (dotted_name) is preceded
	// by "from", not "import"/"," -> returns ""
	src := []byte("from fastapi import FastAPI\n")
	root, _ := parsePython(src)
	stmt := findAllByType(root, "import_from_statement")[0]
	for i := 1; i < int(stmt.ChildCount()); i++ {
		child := stmt.Child(i)
		if child.Type() == "dotted_name" && nodeText(child, src) == "fastapi" {
			if got := tryImportedName(child, stmt, i, src); got != "" {
				t.Fatalf("module name should not be returned, got %q", got)
			}
			return
		}
	}
	t.Skip("module dotted_name not found")
}
