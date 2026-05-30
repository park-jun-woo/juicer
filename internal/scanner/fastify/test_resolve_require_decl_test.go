//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestResolveRequireDecl 테스트
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
