//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestAddImportNames_NoStringChild 테스트
package nestjs

import "testing"

func TestAddImportNames_NoStringChild(t *testing.T) {
	// Pass a non-import node so source child is nil
	src := []byte(`const x = 1;`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	result := make(map[string]string)
	// Use the first child which is not an import_statement
	child := root.Child(0)
	if child != nil {
		addImportNames(child, src, result)
	}
	if len(result) != 0 {
		t.Fatal("expected empty result")
	}
}
