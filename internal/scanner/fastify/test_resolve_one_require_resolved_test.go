//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveOneRequire_Resolved 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveOneRequire_Resolved(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "mod.ts"), []byte(""), 0o644)
	d, fi := reqDeclarator(t, `const m = require("./mod");`+"\n")
	imports := map[string]string{}
	resolveOneRequire(d, fi.Src, dir, imports)
	if imports["m"] == "" {
		t.Fatalf("expected m resolved, got %v", imports)
	}
}
