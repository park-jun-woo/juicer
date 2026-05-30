//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestCollectPrefixCandidates 테스트
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestCollectPrefixCandidates(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main.ts", "x")
	writeFile(t, dir, "src/app.module.ts", "x")
	cands := collectPrefixCandidates(dir)
	if len(cands) < 2 {
		t.Fatalf("expected main.ts + others, got %v", cands)
	}
	if filepath.Base(cands[0]) != "main.ts" {
		t.Fatalf("main.ts should be first: %v", cands)
	}
}
