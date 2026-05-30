//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what resolveRequireDecl 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveRequireDecl(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "mod.ts"), []byte(""), 0o644)
	fi := mustParse(t, []byte(`const m = require("./mod");`+"\n"))
	decls := findAllByType(fi.Root, "lexical_declaration")
	if len(decls) == 0 {
		t.Fatal("no lexical_declaration")
	}
	imports := map[string]string{}
	for _, d := range decls {
		resolveRequireDecl(d, fi.Src, dir, imports)
	}
	if imports["m"] == "" {
		t.Fatalf("expected m resolved, got %v", imports)
	}
}

func TestResolveRequireDecl_NoRequire(t *testing.T) {
	fi := mustParse(t, []byte("const x = 5;\n"))
	decls := findAllByType(fi.Root, "lexical_declaration")
	imports := map[string]string{}
	for _, d := range decls {
		resolveRequireDecl(d, fi.Src, "/d", imports)
	}
	if len(imports) != 0 {
		t.Fatalf("expected no imports, got %v", imports)
	}
}
