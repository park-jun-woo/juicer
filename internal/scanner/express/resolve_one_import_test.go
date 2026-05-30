//ff:func feature=scan type=test control=sequence topic=express
//ff:what resolveOneImport: 상대경로 / alias / 미해결 / 빈경로 분기
package express

import (
	"path/filepath"
	"testing"
)

func TestResolveOneImport_Relative(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "users.ts", "x")
	fi := mustParse(t, []byte(`import r from './users';`))
	imports := map[string]string{}
	resolveOneImport(firstImportStmt(t, fi), fi.Src, dir, imports, dir, nil)
	if imports["r"] != filepath.Join(dir, "users.ts") {
		t.Fatalf("got %v", imports)
	}
}

func TestResolveOneImport_Alias(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/users.ts", "x")
	fi := mustParse(t, []byte(`import r from '@app/users';`))
	imports := map[string]string{}
	aliases := map[string]string{"@app/": filepath.Join("src") + string(filepath.Separator)}
	resolveOneImport(firstImportStmt(t, fi), fi.Src, dir, imports, dir, aliases)
	if imports["r"] == "" {
		t.Fatalf("expected alias resolution, got %v", imports)
	}
}

func TestResolveOneImport_Unresolved(t *testing.T) {
	dir := t.TempDir()
	fi := mustParse(t, []byte(`import r from 'external-pkg';`))
	imports := map[string]string{}
	resolveOneImport(firstImportStmt(t, fi), fi.Src, dir, imports, dir, nil)
	if len(imports) != 0 {
		t.Fatalf("expected unresolved, got %v", imports)
	}
}

func TestResolveOneImport_EmptyPath(t *testing.T) {
	// a node without a string child -> extractImportPath returns "" -> early return
	fi := mustParse(t, []byte(`const x = 1;`))
	imports := map[string]string{}
	resolveOneImport(fi.Root, fi.Src, t.TempDir(), imports, t.TempDir(), nil)
	if len(imports) != 0 {
		t.Fatalf("expected no bindings, got %v", imports)
	}
}
