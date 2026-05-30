//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestParseFile_Success 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile_Success(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "app.ts")
	os.WriteFile(p, []byte("const app = Fastify();\n"), 0o644)
	fi, err := parseFile(p)
	if err != nil {
		t.Fatal(err)
	}
	if fi.Path != p || fi.Root == nil {
		t.Fatalf("unexpected fileInfo: %+v", fi)
	}
}
