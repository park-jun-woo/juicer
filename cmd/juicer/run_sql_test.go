//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_DefaultDirCov 테스트
package main

import (
	"os"
	"testing"
)

func TestRunSQL_DefaultDirCov(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	runSQL([]string{})
}
