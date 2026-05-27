//ff:func feature=ddl type=test control=sequence
//ff:what TestRunDDL_DefaultDirBranch 테스트
package main

import (
	"os"
	"testing"
)

func TestRunDDL_DefaultDirBranch(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	runDDL([]string{})
}
