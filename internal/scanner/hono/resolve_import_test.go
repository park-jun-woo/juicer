//ff:func feature=scan type=test control=sequence topic=hono
//ff:what resolveImports / resolveOneImport 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestResolveImports_Named(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "schemas.ts", "export const s = 1\n")
	writeFile(t, dir, "app.ts", `import { s } from "./schemas"`+"\n")
	fi, err := parseFile(filepath.Join(dir, "app.ts"))
	if err != nil {
		t.Fatal(err)
	}
	imports := resolveImports(fi, dir)
	if imports["s"] != filepath.Join(dir, "schemas.ts") {
		t.Fatalf("got %v", imports)
	}
}

func TestResolveImports_DefaultImport(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "mod.ts", "export default 1\n")
	writeFile(t, dir, "app.ts", `import mod from "./mod"`+"\n")
	fi, _ := parseFile(filepath.Join(dir, "app.ts"))
	imports := resolveImports(fi, dir)
	if imports["mod"] != filepath.Join(dir, "mod.ts") {
		t.Fatalf("got %v", imports)
	}
}

func TestResolveImports_ExternalSkipped(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app.ts", `import { Hono } from "hono"`+"\n")
	fi, _ := parseFile(filepath.Join(dir, "app.ts"))
	if imports := resolveImports(fi, dir); len(imports) != 0 {
		t.Fatalf("expected no imports, got %v", imports)
	}
}

func TestResolveImports_UnresolvableRelative(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app.ts", `import { x } from "./missing"`+"\n")
	fi, _ := parseFile(filepath.Join(dir, "app.ts"))
	if imports := resolveImports(fi, dir); len(imports) != 0 {
		t.Fatalf("expected no imports, got %v", imports)
	}
}
