//ff:func feature=scan type=test control=sequence topic=express
//ff:what default import 변수명 추출 회귀 테스트: import usersRouter from "./routes" → usersRouter
package express

import "testing"

func TestExtractImportVarNameDefault(t *testing.T) {
	src := []byte(`import usersRouter from "./routes";`)
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
