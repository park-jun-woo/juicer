//ff:func feature=scan type=test control=sequence topic=express
//ff:what resolveOneRequire: require해석 / 비require / 비call / 미해결 분기
package express

import (
	"path/filepath"
	"testing"
)

func TestResolveOneRequire_Resolves(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "users.ts", "x")
	fi := mustParse(t, []byte(`const r = require('./users');`))
	imports := map[string]string{}
	resolveOneRequire(firstLexDecl(t, fi), fi.Src, dir, imports, dir, nil)
	if imports["r"] != filepath.Join(dir, "users.ts") {
		t.Fatalf("got %v", imports)
	}
}

func TestResolveOneRequire_NoNameIdentifier(t *testing.T) {
	// destructuring declarator -> no direct identifier name child
	fi := mustParse(t, []byte(`const { a } = require('./users');`))
	imports := map[string]string{}
	resolveOneRequire(firstLexDecl(t, fi), fi.Src, t.TempDir(), imports, t.TempDir(), nil)
	if len(imports) != 0 {
		t.Fatalf("got %v", imports)
	}
}

func TestResolveOneRequire_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`const r = 5;`))
	imports := map[string]string{}
	resolveOneRequire(firstLexDecl(t, fi), fi.Src, t.TempDir(), imports, t.TempDir(), nil)
	if len(imports) != 0 {
		t.Fatalf("got %v", imports)
	}
}

func TestResolveOneRequire_NotRequire(t *testing.T) {
	fi := mustParse(t, []byte(`const r = foo('./x');`))
	imports := map[string]string{}
	resolveOneRequire(firstLexDecl(t, fi), fi.Src, t.TempDir(), imports, t.TempDir(), nil)
	if len(imports) != 0 {
		t.Fatalf("got %v", imports)
	}
}

func TestResolveOneRequire_Unresolved(t *testing.T) {
	fi := mustParse(t, []byte(`const r = require('external-pkg');`))
	imports := map[string]string{}
	resolveOneRequire(firstLexDecl(t, fi), fi.Src, t.TempDir(), imports, t.TempDir(), nil)
	if len(imports) != 0 {
		t.Fatalf("got %v", imports)
	}
}

func TestResolveOneRequire_EmptyPath(t *testing.T) {
	// require with non-string arg -> extractRequirePath empty
	fi := mustParse(t, []byte(`const r = require(modVar);`))
	imports := map[string]string{}
	resolveOneRequire(firstLexDecl(t, fi), fi.Src, t.TempDir(), imports, t.TempDir(), nil)
	if len(imports) != 0 {
		t.Fatalf("got %v", imports)
	}
}
