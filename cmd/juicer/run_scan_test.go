//ff:func feature=scan type=command control=sequence
//ff:what TestRunScan_Cov 테스트
package main

import (
	"path/filepath"
	"testing"
)

func TestRunScan_WithRootCov(t *testing.T) {
	dir := setupMinimalGoProject(t)
	outFile := filepath.Join(dir, "out.yaml")
	runScan([]string{dir, "-o", outFile})
}
