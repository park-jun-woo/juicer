//ff:func feature=ddl type=parse control=sequence
//ff:what TestParse_GlobError 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParse_GlobError(t *testing.T) {
	// filepath.Glob returns ErrBadPattern for malformed patterns
	// A directory name containing '[' without ']' causes this
	base := t.TempDir()
	badDir := filepath.Join(base, "bad[dir")
	if err := os.MkdirAll(badDir, 0o755); err != nil {
		t.Fatal(err)
	}
	_, err := Parse(badDir)
	if err == nil {
		t.Fatal("Parse() expected error for bad glob pattern, got nil")
	}
}
