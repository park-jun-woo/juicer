//ff:func feature=prisma type=test control=sequence
//ff:what runPrisma stdout 렌더링/디렉터리 기록/파싱 에러 분기 직접 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunPrismaDirect(t *testing.T) {
	dir := t.TempDir()
	schema := filepath.Join(dir, "schema.prisma")
	if err := os.WriteFile(schema, []byte("model A { id Int @id }"), 0o644); err != nil {
		t.Fatal(err)
	}
	// stdout render path (outDir == "")
	if err := runPrisma(schema, ""); err != nil {
		t.Errorf("stdout: %v", err)
	}
	// write-files path (outDir != "")
	out := filepath.Join(dir, "out")
	if err := runPrisma(schema, out); err != nil {
		t.Errorf("outdir: %v", err)
	}
	if entries, _ := os.ReadDir(out); len(entries) == 0 {
		t.Error("no files written")
	}
	// parse error path
	if err := runPrisma(filepath.Join(dir, "missing.prisma"), ""); err == nil {
		t.Error("missing schema should error")
	}
}
