//ff:func feature=scan type=test control=iteration dimension=1 topic=supafunc
//ff:what TestFindFunctions 함수 디렉토리 수집 + _shared 스킵 테스트
package supafunc

import (
	"path/filepath"
	"testing"
)

func TestFindFunctions(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "supabase/functions/hello/index.ts", "Deno.serve(async (req) => {})")
	writeFile(t, dir, "supabase/functions/get-user/index.ts", "serve(async (req) => {})")
	writeFile(t, dir, "supabase/functions/_shared/cors.ts", "export const corsHeaders = {}")

	files, err := findFunctions(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 2 {
		t.Fatalf("expected 2 functions, got %d", len(files))
	}
	for _, f := range files {
		base := filepath.Base(filepath.Dir(f))
		if base == "_shared" {
			t.Fatalf("_shared should be skipped, got %s", f)
		}
	}
}
