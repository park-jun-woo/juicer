//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_WithRootCov 테스트
package main

import (
	"path/filepath"
	"testing"
)

func TestRunScan_WithRootCov(t *testing.T) {
	dir := setupMinimalGoProject(t)
	outFile := filepath.Join(dir, "out.yaml")
	execScan([]string{"-o", outFile, dir})
}
