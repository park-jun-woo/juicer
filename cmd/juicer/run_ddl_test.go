//ff:func feature=ddl type=command control=sequence
//ff:what TestRunDDL_Cov 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunDDL_StdoutCov(t *testing.T) {
	dir := t.TempDir()
	sql := "CREATE TABLE users (id INT PRIMARY KEY, name TEXT);\n"
	os.WriteFile(filepath.Join(dir, "001.sql"), []byte(sql), 0o644)
	runDDL([]string{dir})
}

func TestRunDDL_ToDirCov(t *testing.T) {
	dir := t.TempDir()
	outDir := filepath.Join(dir, "out")
	sql := "CREATE TABLE users (id INT PRIMARY KEY, name TEXT);\n"
	os.WriteFile(filepath.Join(dir, "001.sql"), []byte(sql), 0o644)
	runDDL([]string{dir, "-o", outDir})
}

func TestRunDDL_EmptyDirCov(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	runDDL([]string{})
}
