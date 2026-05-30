//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestScanPass2_SkipFilesWithoutVars 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestScanPass2_SkipFilesWithoutVars(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app.ts", scanInternalsSrc)
	writeFile(t, dir, "plain.ts", "const x = 1\n")
	ctx := scanPass1([]string{
		filepath.Join(dir, "app.ts"),
		filepath.Join(dir, "plain.ts"),
	}, dir)
	eps := scanPass2(ctx)

	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(eps))
	}
}
