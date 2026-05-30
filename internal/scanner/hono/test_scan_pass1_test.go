//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestScanPass1 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestScanPass1(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app.ts", scanInternalsSrc)
	ctx := scanPass1([]string{filepath.Join(dir, "app.ts")}, dir)
	if ctx == nil || len(ctx.parsed) != 1 {
		t.Fatalf("unexpected ctx: %+v", ctx)
	}
	if ctx.absRoot != dir {
		t.Fatalf("absRoot: %s", ctx.absRoot)
	}
}
