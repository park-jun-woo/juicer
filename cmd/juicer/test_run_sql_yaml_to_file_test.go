//ff:func feature=scan type=command control=sequence
//ff:what TestRunSQL_YAML_ToFile 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunSQL_YAML_ToFile(t *testing.T) {
	dir := t.TempDir()
	outFile := filepath.Join(dir, "output.yaml")
	runSQL([]string{"-o", outFile, dir})

	if _, err := os.Stat(outFile); err != nil {
		t.Errorf("expected output file: %v", err)
	}
}
