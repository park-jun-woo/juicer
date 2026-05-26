//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestAddImportNames_NoImportClause 테스트
package nestjs

import "testing"

func TestAddImportNames_NoImportClause(t *testing.T) {
	// Side-effect-only import has no import_clause
	src := []byte(`import './polyfills';`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	result := make(map[string]string)
	stmt := findChildByType(root, "import_statement")
	if stmt != nil {
		addImportNames(stmt, src, result)
	}
	if len(result) != 0 {
		t.Fatal("expected empty result for side-effect import")
	}
}
