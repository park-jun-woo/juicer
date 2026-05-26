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
	runSQL([]string{dir})

	// json stdout
	runSQL([]string{"-json", dir})

	// output to file
	outFile := filepath.Join(t.TempDir(), "out.yaml")
	runSQL([]string{"-o", outFile, dir})

	// default dir = "."
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	runSQL([]string{})
}
