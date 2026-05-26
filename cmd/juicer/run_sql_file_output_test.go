//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_FileOutput 파일 출력 분기 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunSQL_FileOutput(t *testing.T) {
	dir := t.TempDir()
	out := filepath.Join(t.TempDir(), "out.yaml")
	runSQL([]string{"-o", out, dir})
	if _, err := os.Stat(out); err != nil {
		t.Fatal("expected output file")
	}
}
