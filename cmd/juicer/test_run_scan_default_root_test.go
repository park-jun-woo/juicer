//ff:func feature=scan type=command control=sequence
//ff:what TestRunScan_DefaultRoot 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunScan_DefaultRoot(t *testing.T) {
	// Test without specifying root directory (uses ".")
	// This tests the default root = "." path
	// Change to a temp dir to avoid scanning the whole project
	dir := setupMinimalGoProject(t)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	outFile := filepath.Join(dir, "out.yaml")
	runScan([]string{"-o", outFile})
}
