//ff:func feature=scan type=test topic=express control=sequence
//ff:what loadParsedFile 캐시 적중/파싱/실패(nil) 분기 테스트
package express

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadParsedFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "r.ts")
	if err := os.WriteFile(path, []byte(`const x = 1;`), 0o644); err != nil {
		t.Fatal(err)
	}
	ctx := &scanContext{parsed: map[string]*fileInfo{}}
	// first load parses and caches
	fi := loadParsedFile(ctx, path)
	if fi == nil {
		t.Fatal("expected parsed fileInfo")
	}
	if ctx.parsed[path] != fi {
		t.Error("not cached")
	}
	// cache hit returns same pointer
	if loadParsedFile(ctx, path) != fi {
		t.Error("cache hit should return same fi")
	}
	// parse error -> nil
	if loadParsedFile(ctx, filepath.Join(dir, "missing.ts")) != nil {
		t.Error("missing file should return nil")
	}
}
