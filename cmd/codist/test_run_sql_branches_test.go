//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_Branches 서브커맨드 전체 분기 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunSQL_Branches(t *testing.T) {
	dir := t.TempDir()

	// yaml stdout (default, empty dir)
	execSQL([]string{dir})

	// json stdout
	execSQL([]string{"--json", dir})

	// output to file
	outFile := filepath.Join(t.TempDir(), "out.yaml")
	execSQL([]string{"-o", outFile, dir})

	// default dir = "."
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	execSQL([]string{})
}
