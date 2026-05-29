//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_NoArgs 인자 없음 분기 테스트
package main

import (
	"os"
	"testing"
)

func TestRunSQL_NoArgs(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	execSQL([]string{})
}
