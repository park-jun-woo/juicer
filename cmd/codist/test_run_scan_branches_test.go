//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_Branches 출력 분기 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunScan_Branches(t *testing.T) {
	dir := setupMinimalGoProject(t)

	// yaml stdout (default)
	runScan([]string{dir})

	// json stdout
	runScan([]string{"-json", dir})

	// openapi stdout
	runScan([]string{"-openapi", dir})

	// output to file
	outFile := filepath.Join(t.TempDir(), "out.yaml")
	runScan([]string{"-o", outFile, dir})

	// default root = "."
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	runScan([]string{})
}
