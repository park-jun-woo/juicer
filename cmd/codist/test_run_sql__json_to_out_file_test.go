//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_JSONToOutFile 테스트
package main

import (
	"path/filepath"
	"testing"
)

func TestRunSQL_JSONToOutFile(t *testing.T) {
	dir := t.TempDir()
	outFile := filepath.Join(dir, "out.json")
	execSQL([]string{"--json", "-o", outFile, dir})
}
