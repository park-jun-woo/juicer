//ff:func feature=prisma type=test control=sequence topic=prisma
//ff:what Parse 단일 schema.prisma 파일 변환 및 에러 경로 테스트
package prisma

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseDirect(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "schema.prisma")
	src := `enum Role {
  ADMIN
  USER
}
model User {
  id Int @id @default(autoincrement())
  role Role
}`
	if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	tables, enums, err := Parse(path)
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	if len(tables) != 1 {
		t.Errorf("tables: %v", tables)
	}
	if len(enums) != 1 || enums[0].Name != `"Role"` {
		t.Errorf("enums: %+v", enums)
	}
	// collect error: missing path
	if _, _, err := Parse(filepath.Join(dir, "missing")); err == nil {
		t.Error("missing path should error")
	}
}
