//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestTryImportedName_DottedName 테스트
package fastapi

import "testing"

func TestTryImportedName_DottedName(t *testing.T) {

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
