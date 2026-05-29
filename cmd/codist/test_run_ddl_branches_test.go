//ff:func feature=ddl type=test control=sequence
//ff:what TestRunDDL_Branches 서브커맨드 분기 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunDDL_Branches(t *testing.T) {
	// default dir (empty)
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	execDDL([]string{})

	// explicit dir with render to stdout
	srcDir := t.TempDir()
	sql := "CREATE TABLE items (id INT PRIMARY KEY, name TEXT);\n"
	os.WriteFile(filepath.Join(srcDir, "001.sql"), []byte(sql), 0o644)
	execDDL([]string{srcDir})

	// explicit dir with -o output dir
	outDir := filepath.Join(t.TempDir(), "out")
	execDDL([]string{"-o", outDir, srcDir})
}
