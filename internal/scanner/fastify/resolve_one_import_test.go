//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what resolveOneImport 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveOneImport_Resolved(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "mod.ts"), []byte(""), 0o644)
	fi := mustParse(t, []byte(`import m from "./mod";`+"\n"))
	stmt := findAllByType(fi.Root, "import_statement")[0]
	imports := map[string]string{}
	resolveOneImport(stmt, fi.Src, dir, imports)
	if imports["m"] == "" {
		t.Fatalf("expected m resolved, got %v", imports)
	}
}

func TestResolveOneImport_External(t *testing.T) {
	// external module path doesn't resolve to a local file -> not added
	dir := t.TempDir()
	fi := mustParse(t, []byte(`import Fastify from "fastify";`+"\n"))
	stmt := findAllByType(fi.Root, "import_statement")[0]
	imports := map[string]string{}
	resolveOneImport(stmt, fi.Src, dir, imports)
	if len(imports) != 0 {
		t.Fatalf("expected no imports for external module, got %v", imports)
	}
}

func TestResolveOneImport_SideEffect(t *testing.T) {
	// side-effect import has no var name -> early return
	dir := t.TempDir()
	fi := mustParse(t, []byte(`import "./styles.css";`+"\n"))
	stmt := findAllByType(fi.Root, "import_statement")[0]
	imports := map[string]string{}
	resolveOneImport(stmt, fi.Src, dir, imports)
	if len(imports) != 0 {
		t.Fatalf("expected no imports for side-effect, got %v", imports)
	}
}
