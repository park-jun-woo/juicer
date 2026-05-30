//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what TestTryImportedName_NotPrecededByImport 테스트
package fastapi

import "testing"

func TestTryImportedName_NotPrecededByImport(t *testing.T) {

	src := []byte("from fastapi import FastAPI\n")
	root, _ := parsePython(src)
	stmt := findAllByType(root, "import_from_statement")[0]
	for i := 1; i < int(stmt.ChildCount()); i++ {
		child := stmt.Child(i)
		if child.Type() != "dotted_name" || nodeText(child, src) != "fastapi" {
			continue
		}
		if got := tryImportedName(child, stmt, i, src); got != "" {
			t.Fatalf("module name should not be returned, got %q", got)
		}
		return
	}
	t.Skip("module dotted_name not found")
}
