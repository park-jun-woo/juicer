//ff:func feature=scan type=command control=sequence
//ff:what TestRunScan_JSON 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunScan_JSON(t *testing.T) {
	dir := setupMinimalGoProject(t)
	outFile := filepath.Join(dir, "output.json")
	runScan([]string{"-json", "-o", outFile, dir})

	if _, err := os.Stat(outFile); err != nil {
		t.Fatalf("expected output file: %v", err)
	}
}
