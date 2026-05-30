//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what scanPass1 테스트
package fastify

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanPass1(t *testing.T) {
	dir := t.TempDir()
	p1 := filepath.Join(dir, "a.ts")
	os.WriteFile(p1, []byte("const app = Fastify();\napp.get('/x', h);\n"), 0o644)
	p2 := filepath.Join(dir, "b.ts")
	os.WriteFile(p2, []byte("const srv = fastify();\n"), 0o644)

	// include one unreadable path -> scanOneFilePass1 returns nil -> continue
	ctx := scanPass1([]string{p1, p2, "/no/such.ts"}, dir)
	if len(ctx.parsed) != 2 {
		t.Fatalf("expected 2 parsed files, got %d", len(ctx.parsed))
	}
	if !ctx.instances[p1]["app"] {
		t.Errorf("expected app instance in a.ts")
	}
	if ctx.absRoot != dir {
		t.Errorf("absRoot = %q", ctx.absRoot)
	}
}
