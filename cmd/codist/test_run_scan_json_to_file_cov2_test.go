//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_JSONToFileCov2 테스트
package main

import (
	"path/filepath"
	"testing"
)

func TestRunScan_JSONToFileCov2(t *testing.T) {
	dir := setupMinimalGoProject(t)
	outFile := filepath.Join(dir, "out.json")
	execScan([]string{"--json", "-o", outFile, dir})
}
