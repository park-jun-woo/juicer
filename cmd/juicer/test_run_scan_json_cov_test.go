//ff:func feature=scan type=command control=sequence
//ff:what TestRunScan_JSONCov 테스트
package main

import (
	"path/filepath"
	"testing"
)

func TestRunScan_JSONCov(t *testing.T) {
	dir := setupMinimalGoProject(t)
	outFile := filepath.Join(dir, "out.json")
	runScan([]string{dir, "--json", "-o", outFile})
}
