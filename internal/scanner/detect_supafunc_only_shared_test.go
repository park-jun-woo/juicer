//ff:func feature=scan type=test control=sequence
//ff:what TestDetectSupaFunc_OnlyShared _shared만 있을 때 감지 실패 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectSupaFunc_OnlyShared(t *testing.T) {
	dir := t.TempDir()
	sharedDir := filepath.Join(dir, "supabase", "functions", "_shared")
	os.MkdirAll(sharedDir, 0o755)
	os.WriteFile(filepath.Join(sharedDir, "cors.ts"), []byte("export const x = 1"), 0o644)
	if detectSupaFunc(dir) {
		t.Fatal("expected false when only _shared exists")
	}
}
