//ff:func feature=scan type=command control=sequence
//ff:what TestRunScan_YAML_ToFile 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunScan_YAML_ToFile(t *testing.T) {
	dir := setupMinimalGoProject(t)
	outFile := filepath.Join(dir, "output.yaml")
	execScan([]string{"-o", outFile, dir})

	if _, err := os.Stat(outFile); err != nil {
		t.Fatalf("expected output file: %v", err)
	}
}
