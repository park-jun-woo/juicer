//ff:func feature=ddl type=test control=sequence
//ff:what runDDL stdout/outDir/파싱에러 경로 직접 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunDDLAllPaths(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "001.sql"), []byte("CREATE TABLE foo (id integer);"), 0o644); err != nil {
		t.Fatal(err)
	}
	// stdout
	if err := runDDL(dir, ""); err != nil {
		t.Errorf("stdout: %v", err)
	}
	// outDir
	out := filepath.Join(dir, "out")
	if err := runDDL(dir, out); err != nil {
		t.Errorf("outdir: %v", err)
	}
	// write error: outDir is an existing file, not a dir
	badOut := filepath.Join(dir, "001.sql")
	if err := runDDL(dir, badOut); err == nil {
		t.Error("outDir as file should error")
	}
}
