//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestScanOneFilePass1_OK 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestScanOneFilePass1_OK(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app.ts", scanInternalsSrc)
	r := scanOneFilePass1(filepath.Join(dir, "app.ts"), dir)
	if r == nil {
		t.Fatal("nil result")
	}
	if r.fi == nil || len(r.vars) == 0 {
		t.Fatalf("missing fi/vars: %+v", r)
	}
}
