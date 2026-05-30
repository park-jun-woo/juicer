//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestScanPass2 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestScanPass2(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app.ts", scanInternalsSrc)
	ctx := scanPass1([]string{filepath.Join(dir, "app.ts")}, dir)
	eps := scanPass2(ctx)
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d: %+v", len(eps), eps)
	}
}
