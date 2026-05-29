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
	execScan([]string{dir})

	// json stdout
	execScan([]string{"--json", dir})

	// openapi stdout
	execScan([]string{"--openapi", dir})

	// output to file
	outFile := filepath.Join(t.TempDir(), "out.yaml")
	execScan([]string{"-o", outFile, dir})

	// default root = "."
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	execScan([]string{})
}
