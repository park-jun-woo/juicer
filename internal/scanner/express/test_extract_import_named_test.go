//ff:func feature=scan type=test control=sequence topic=express
//ff:what named import 변수명 추출 테스트: import { healthCheckRouter } from "./routes" → healthCheckRouter
package express

import "testing"

func TestExtractImportVarNameNamed(t *testing.T) {
	src := []byte(`import { healthCheckRouter } from "./routes";`)
	fi := mustParse(t, src)
	stmts := findAllByType(fi.Root, "import_statement")
	if len(stmts) != 1 {
		t.Fatalf("expected 1 import_statement, got %d", len(stmts))
	}
	got := extractImportVarName(stmts[0], src)
	if got != "healthCheckRouter" {
		t.Errorf("want healthCheckRouter, got %s", got)
	}
}
