//ff:func feature=ddl type=test control=sequence
//ff:what TestRunDDL_OutDirBranch 출력 디렉터리 분기 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunDDL_OutDirBranch(t *testing.T) {
	dir := t.TempDir()
	sql := "CREATE TABLE users (id INT PRIMARY KEY, name TEXT);\n"
	os.WriteFile(filepath.Join(dir, "001.sql"), []byte(sql), 0o644)
	outDir := filepath.Join(t.TempDir(), "out")
	runDDL([]string{"-o", outDir, dir})
}
