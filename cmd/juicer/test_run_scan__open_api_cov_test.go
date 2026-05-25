//ff:func feature=scan type=command control=sequence
//ff:what TestRunScan_OpenAPICov 테스트
package main

import (
	"path/filepath"
	"testing"
)

func TestRunScan_OpenAPICov(t *testing.T) {
	dir := setupMinimalGoProject(t)
	outFile := filepath.Join(dir, "out.yaml")
	runScan([]string{dir, "--openapi", "-o", outFile})
}
