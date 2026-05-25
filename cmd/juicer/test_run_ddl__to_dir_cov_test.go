//ff:func feature=ddl type=command control=sequence
//ff:what TestRunDDL_ToDirCov 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunDDL_ToDirCov(t *testing.T) {
	dir := t.TempDir()
	outDir := filepath.Join(dir, "out")
	sql := "CREATE TABLE users (id INT PRIMARY KEY, name TEXT);\n"
	os.WriteFile(filepath.Join(dir, "001.sql"), []byte(sql), 0o644)
	runDDL([]string{dir, "-o", outDir})
}
