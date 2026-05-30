//ff:func feature=scan type=test control=sequence topic=express
//ff:what collectImportedNames — import 이름 수집 분기를 검증
package express

import "testing"

func sortedHas(names []string, want string) bool {
	for _, n := range names {
		if n == want {
			return true
		}
	}
	return false
}

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

func TestCollectImportedNames_NamedOnly(t *testing.T) {
	fi := mustParse(t, []byte("import { Router } from 'express';\n"))
	stmt := firstParamOfType(fi.Root, "import_statement")
	names := collectImportedNames(stmt, fi.Src)
	if !sortedHas(names, "Router") {
		t.Fatalf("expected Router, got %v", names)
	}
}

func TestCollectImportedNames_DefaultOnly(t *testing.T) {
	fi := mustParse(t, []byte("import express from 'express';\n"))
	stmt := firstParamOfType(fi.Root, "import_statement")
	names := collectImportedNames(stmt, fi.Src)
	if len(names) != 1 || names[0] != "express" {
		t.Fatalf("expected [express], got %v", names)
	}
}

func TestCollectImportedNames_NoClause(t *testing.T) {
	// `import 'side-effect';` has no import_clause.
	fi := mustParse(t, []byte("import 'side-effect';\n"))
	stmt := firstParamOfType(fi.Root, "import_statement")
	if stmt == nil {
		t.Fatal("no import_statement")
	}
	if names := collectImportedNames(stmt, fi.Src); names != nil {
		t.Fatalf("expected nil for side-effect import, got %v", names)
	}
}
