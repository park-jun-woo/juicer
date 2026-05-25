//ff:func feature=scan type=command control=sequence
//ff:what TestRunDDL_DefaultDir 테스트
package main

import (
	"os"
	"testing"
)

func TestRunDDL_DefaultDir(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	runDDL([]string{})
}
