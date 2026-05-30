//ff:func feature=scan type=test control=sequence topic=hono
//ff:what resolveOneImport 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func resolveOne(t *testing.T, dir, appSrc string) map[string]string {
	t.Helper()
	writeFile(t, dir, "app.ts", appSrc+"\n")
	fi, err := parseFile(filepath.Join(dir, "app.ts"))
	if err != nil {
		t.Fatal(err)
	}
	imports := map[string]string{}
	for _, stmt := range findAllByType(fi.Root, "import_statement") {
		resolveOneImport(stmt, fi.Src, dir, imports, dir)
	}
	return imports
}

func TestResolveOneImport_NoPathNode(t *testing.T) {
	dir := t.TempDir()
	// bare import side-effect with no module string is rare; use a non-import stmt scenario
	imp := resolveOne(t, dir, `const x = 1`)
	if len(imp) != 0 {
		t.Fatalf("expected none, got %v", imp)
	}
}

func TestResolveOneImport_External(t *testing.T) {
	dir := t.TempDir()
	imp := resolveOne(t, dir, `import { z } from "zod"`)
	if len(imp) != 0 {
		t.Fatalf("expected external skipped, got %v", imp)
	}
}

func TestResolveOneImport_Unresolvable(t *testing.T) {
	dir := t.TempDir()
	imp := resolveOne(t, dir, `import { x } from "./nope"`)
	if len(imp) != 0 {
		t.Fatalf("expected unresolved skipped, got %v", imp)
	}
}

func TestResolveOneImport_Named(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "schemas.ts", "export const a = 1\n")
	imp := resolveOne(t, dir, `import { a } from "./schemas"`)
	if imp["a"] != filepath.Join(dir, "schemas.ts") {
		t.Fatalf("got %v", imp)
	}
}

func TestResolveOneImport_Default(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "mod.ts", "export default 1\n")
	imp := resolveOne(t, dir, `import mod from "./mod"`)
	if imp["mod"] != filepath.Join(dir, "mod.ts") {
		t.Fatalf("got %v", imp)
	}
}
