//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectImportedNames_NoClause 테스트
package express

import "testing"

func TestCollectImportedNames_NoClause(t *testing.T) {

	fi := mustParse(t, []byte("import 'side-effect';\n"))
	stmt := firstParamOfType(fi.Root, "import_statement")
	if stmt == nil {
		t.Fatal("no import_statement")
	}
	if names := collectImportedNames(stmt, fi.Src); names != nil {
		t.Fatalf("expected nil for side-effect import, got %v", names)
	}
}
