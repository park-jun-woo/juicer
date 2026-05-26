//ff:func feature=scan type=test control=sequence
//ff:what TestDetectNestJS_AllBranches miss/no-nestjs 전 분기 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectNestJS_AllBranches(t *testing.T) {
	// no package.json
	dir1 := t.TempDir()
	if detectNestJS(dir1) {
		t.Fatal("expected false when no package.json")
	}

	// package.json without nestjs
	dir2 := t.TempDir()
	os.WriteFile(filepath.Join(dir2, "package.json"), []byte(`{"dependencies":{"express":"*"}}`), 0o644)
	if detectNestJS(dir2) {
		t.Fatal("expected false when nestjs not present")
	}
}
