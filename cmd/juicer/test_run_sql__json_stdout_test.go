//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_JSONStdout 테스트
package main

import (
	"os"
	"testing"
)

func TestRunSQL_JSONStdout(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	runSQL([]string{"-json"})
}
