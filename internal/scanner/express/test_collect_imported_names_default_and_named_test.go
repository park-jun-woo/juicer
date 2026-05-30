//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCollectImportedNames_DefaultAndNamed 테스트
package express

import "testing"

func TestCollectImportedNames_DefaultAndNamed(t *testing.T) {
	fi := mustParse(t, []byte("import express, { Router, Express } from 'express';\n"))
	stmt := firstParamOfType(fi.Root, "import_statement")
	if stmt == nil {
		t.Fatal("no import_statement")
	}
	names := collectImportedNames(stmt, fi.Src)
	if !sortedHas(names, "express") || !sortedHas(names, "Router") || !sortedHas(names, "Express") {
		t.Fatalf("expected express, Router, Express; got %v", names)
	}
}
