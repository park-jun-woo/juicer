//ff:func feature=scan type=command control=sequence
//ff:what TestRunSQL_JSON_ToFile 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunSQL_JSON_ToFile(t *testing.T) {
	dir := t.TempDir()
	outFile := filepath.Join(dir, "output.json")
	runSQL([]string{"-json", "-o", outFile, dir})

	if _, err := os.Stat(outFile); err != nil {
		t.Errorf("expected output file: %v", err)
	}
}
