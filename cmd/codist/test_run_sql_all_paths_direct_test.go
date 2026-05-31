//ff:func feature=sql type=test control=sequence
//ff:what runSQL JSON/YAML stdout, outFile, extract 에러, write 에러 경로 직접 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunSQLAllPathsDirect(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	// yaml stdout
	if err := runSQL(dir, false, ""); err != nil {
		t.Errorf("yaml stdout: %v", err)
	}
	// json stdout
	if err := runSQL(dir, true, ""); err != nil {
		t.Errorf("json stdout: %v", err)
	}
	// to file
	out := filepath.Join(dir, "o.yaml")
	if err := runSQL(dir, false, out); err != nil {
		t.Errorf("to file: %v", err)
	}
	// write error: outFile inside nonexistent dir
	if err := runSQL(dir, false, filepath.Join(dir, "nope", "o.yaml")); err == nil {
		t.Error("bad outFile should error")
	}
}
