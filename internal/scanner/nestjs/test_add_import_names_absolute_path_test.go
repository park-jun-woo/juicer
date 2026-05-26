//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestAddImportNames_AbsolutePath 테스트
package nestjs

import "testing"

func TestAddImportNames_AbsolutePath(t *testing.T) {
	src := []byte(`import { Module } from '@nestjs/common';`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	result := make(map[string]string)
	stmt := findChildByType(root, "import_statement")
	if stmt == nil {
		t.Fatal("expected import_statement")
	}
	addImportNames(stmt, src, result)
	if len(result) != 0 {
		t.Fatalf("expected empty result for non-relative path, got %v", result)
	}
}
