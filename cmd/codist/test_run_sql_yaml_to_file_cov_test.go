//ff:func feature=sql type=command control=sequence
//ff:what TestRunSQL_YAMLToFileCov 테스트
package main

import (
	"path/filepath"
	"testing"
)

func TestRunSQL_YAMLToFileCov(t *testing.T) {
	dir := t.TempDir()
	outFile := filepath.Join(dir, "out.yaml")
	runSQL([]string{dir, "-o", outFile})
}
