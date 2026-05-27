//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_WithDirYAMLToFileCov2 테스트
package main

import (
	"path/filepath"
	"testing"
)

func TestRunSQL_WithDirYAMLToFileCov2(t *testing.T) {
	dir := t.TempDir()
	outFile := filepath.Join(dir, "out.yaml")
	runSQL([]string{"-o", outFile, dir})
}
