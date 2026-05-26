//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_DefaultRootDir 테스트
package main

import (
	"os"
	"testing"
)

func TestRunScan_DefaultRootDir(t *testing.T) {
	dir := setupMinimalGoProject(t)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	runScan([]string{})
}
