//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractImportPath_Round5 테스트
package express

import "testing"

func TestExtractImportPath_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`import x from './mod';`))
	stmt := exFirst(t, fi, "import_statement")
	if got := extractImportPath(stmt, fi.Src); got != "./mod" {
		t.Fatalf("got %q", got)
	}
}
