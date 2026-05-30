//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestResolveOneRequire_IndexResolution 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveOneRequire_IndexResolution(t *testing.T) {
	dir := t.TempDir()
	sub := filepath.Join(dir, "pkg")
	os.MkdirAll(sub, 0o755)
	os.WriteFile(filepath.Join(sub, "index.ts"), []byte(""), 0o644)
	d, fi := reqDeclarator(t, `const p = require("./pkg");`+"\n")
	imports := map[string]string{}
	resolveOneRequire(d, fi.Src, dir, imports)
	if imports["p"] == "" {
		t.Fatalf("expected p resolved to index.ts, got %v", imports)
	}
}
