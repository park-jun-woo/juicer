//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what collectPrismaFiles 파일/디렉터리 글롭 및 stat 에러 테스트
package prisma

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCollectPrismaFiles(t *testing.T) {
	dir := t.TempDir()
	a := filepath.Join(dir, "schema.prisma")
	if err := os.WriteFile(a, []byte("model A { id Int @id }"), 0o644); err != nil {
		t.Fatal(err)
	}
	// single file
	files, err := collectPrismaFiles(a)
	if err != nil || len(files) != 1 || files[0] != a {
		t.Errorf("file: %v %v", files, err)
	}
	// directory glob
	files, err = collectPrismaFiles(dir)
	if err != nil || len(files) != 1 {
		t.Errorf("dir: %v %v", files, err)
	}
	// stat error
	if _, err := collectPrismaFiles(filepath.Join(dir, "nope")); err == nil {
		t.Error("missing path should error")
	}
	// glob error: a directory name with an unterminated bracket makes the
	// joined pattern (".../[/*.prisma") malformed -> filepath.ErrBadPattern.
	bad := filepath.Join(dir, "[")
	if err := os.Mkdir(bad, 0o755); err != nil {
		t.Fatal(err)
	}
	if _, err := collectPrismaFiles(bad); err == nil {
		t.Error("malformed glob pattern should error")
	}
}
