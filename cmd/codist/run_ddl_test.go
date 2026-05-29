//ff:func feature=ddl type=test control=sequence
//ff:what TestRunDDL_StdoutBranch stdout 분기 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunDDL_StdoutBranch(t *testing.T) {
	dir := t.TempDir()
	sql := "CREATE TABLE users (id INT PRIMARY KEY, name TEXT);\n"
	if err := os.WriteFile(filepath.Join(dir, "001.sql"), []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}

	// stdout branch: no -o flag, dir passed as positional arg
	execDDL([]string{dir})

	// output-dir branch: -o writes one .sql file per table
	outDir := filepath.Join(t.TempDir(), "out")
	execDDL([]string{"-o", outDir, dir})

	entries, err := os.ReadDir(outDir)
	if err != nil {
		t.Fatalf("output dir not created: %v", err)
	}
	if len(entries) == 0 {
		t.Fatal("expected at least one .sql file written, got none")
	}
}
