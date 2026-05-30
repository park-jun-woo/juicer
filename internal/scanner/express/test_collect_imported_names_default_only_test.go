//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectImportedNames_DefaultOnly 테스트
package express

import "testing"

func TestCollectImportedNames_DefaultOnly(t *testing.T) {
	fi := mustParse(t, []byte("import express from 'express';\n"))
	stmt := firstParamOfType(fi.Root, "import_statement")
	names := collectImportedNames(stmt, fi.Src)
	if len(names) != 1 || names[0] != "express" {
		t.Fatalf("expected [express], got %v", names)
	}
}
