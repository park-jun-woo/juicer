//ff:func feature=ddl type=command control=sequence
//ff:what TestRunDDL_EmptyDirCov 테스트
package main

import (
	"os"
	"testing"
)

func TestRunDDL_EmptyDirCov(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	execDDL([]string{})
}
