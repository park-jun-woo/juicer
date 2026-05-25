//ff:func feature=scan type=command control=sequence
//ff:what TestRunScan_OpenAPI 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunScan_OpenAPI(t *testing.T) {
	dir := setupMinimalGoProject(t)
	outFile := filepath.Join(dir, "output.openapi.yaml")
	runScan([]string{"-openapi", "-o", outFile, dir})

	if _, err := os.Stat(outFile); err != nil {
		t.Fatalf("expected output file: %v", err)
	}
}
