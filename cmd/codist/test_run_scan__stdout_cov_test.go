//ff:func feature=scan type=command control=sequence
//ff:what TestRunScan_StdoutCov 테스트
package main

import (
	"os"
	"testing"
)

func TestRunScan_StdoutCov(t *testing.T) {
	dir := setupMinimalGoProject(t)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	execScan([]string{})
}
