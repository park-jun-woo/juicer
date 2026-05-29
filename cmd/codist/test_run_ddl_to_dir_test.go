//ff:func feature=scan type=command control=sequence
//ff:what TestRunDDL_ToDir 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunDDL_ToDir(t *testing.T) {
	dir := t.TempDir()
	sql := `CREATE TABLE users (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL
);
`
	os.WriteFile(filepath.Join(dir, "001_init.up.sql"), []byte(sql), 0o644)
	outDir := filepath.Join(dir, "output")
	execDDL([]string{"-o", outDir, dir})

	matches, _ := filepath.Glob(filepath.Join(outDir, "*_users.sql"))
	if len(matches) == 0 {
		t.Errorf("expected a *_users.sql file to exist in %s", outDir)
	}
}
