//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectImportedNames_NamedOnly 테스트
package express

import "testing"

func TestCollectImportedNames_NamedOnly(t *testing.T) {
	fi := mustParse(t, []byte("import { Router } from 'express';\n"))
	stmt := firstParamOfType(fi.Root, "import_statement")
	names := collectImportedNames(stmt, fi.Src)
	if !sortedHas(names, "Router") {
		t.Fatalf("expected Router, got %v", names)
	}
}
