//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_WithDirAndJSONCov2 테스트
package main

import (
	"path/filepath"
	"testing"
)

func TestRunSQL_WithDirAndJSONCov2(t *testing.T) {
	dir := t.TempDir()
	outFile := filepath.Join(dir, "out.json")
	execSQL([]string{"--json", "-o", outFile, dir})
}
