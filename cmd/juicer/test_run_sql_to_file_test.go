//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_ToFile 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunSQL_ToFile(t *testing.T) {
	dir := t.TempDir()
	out := filepath.Join(t.TempDir(), "out.yaml")
	runSQL([]string{"-o", out, dir})
	if _, err := os.Stat(out); err != nil {
		t.Fatalf("output file not created: %v", err)
	}
}
