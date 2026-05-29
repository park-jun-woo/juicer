//ff:func feature=ddl type=test control=sequence
//ff:what TestRunDDL_WriteFileBranch 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunDDL_WriteFileBranch(t *testing.T) {
	dir := t.TempDir()
	sql := "CREATE TABLE users (id INT PRIMARY KEY, name TEXT);\n"
	os.WriteFile(filepath.Join(dir, "001.sql"), []byte(sql), 0o644)
	outDir := t.TempDir()
	execDDL([]string{"-o", outDir, dir})
}
