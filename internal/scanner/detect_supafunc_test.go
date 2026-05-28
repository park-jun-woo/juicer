//ff:func feature=scan type=test control=sequence
//ff:what TestDetectSupaFunc_Hit supabase/functions 구조 감지 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectSupaFunc_Hit(t *testing.T) {
	dir := t.TempDir()
	funcDir := filepath.Join(dir, "supabase", "functions", "hello")
	os.MkdirAll(funcDir, 0o755)
	os.WriteFile(filepath.Join(funcDir, "index.ts"), []byte("Deno.serve(async (req) => {})"), 0o644)
	if !detectSupaFunc(dir) {
		t.Fatal("expected true")
	}
}
