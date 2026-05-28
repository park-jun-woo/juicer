//ff:func feature=scan type=test control=sequence topic=express
//ff:what named import alias 추출 테스트: import { router as usersRouter } from "./routes" → usersRouter
package express

import "testing"

func TestExtractImportVarNameAlias(t *testing.T) {
	src := []byte(`import { router as usersRouter } from "./routes";`)
	fi := mustParse(t, src)
	stmts := findAllByType(fi.Root, "import_statement")
	if len(stmts) != 1 {
		t.Fatalf("expected 1 import_statement, got %d", len(stmts))
	}
	got := extractImportVarName(stmts[0], src)
	if got != "usersRouter" {
		t.Errorf("want usersRouter, got %s", got)
	}
}
