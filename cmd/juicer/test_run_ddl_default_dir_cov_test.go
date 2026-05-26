//ff:func feature=ddl type=test control=sequence
//ff:what TestRunDDL_DefaultDirCov 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunDDL_DefaultDirCov(t *testing.T) {
	dir := t.TempDir()
	sql := "CREATE TABLE items (id INT PRIMARY KEY, price DECIMAL);\n"
	os.WriteFile(filepath.Join(dir, "001.sql"), []byte(sql), 0o644)
	outDir := filepath.Join(dir, "out2")
	runDDL([]string{"-o", outDir, dir})
}
