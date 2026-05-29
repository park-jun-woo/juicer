//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_ToFile 파일 출력 분기 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunScan_ToFile(t *testing.T) {
	dir := setupMinimalGoProject(t)
	out := filepath.Join(t.TempDir(), "out.yaml")
	execScan([]string{"-o", out, dir})
	if _, err := os.Stat(out); err != nil {
		t.Fatal("expected output file")
	}
}
