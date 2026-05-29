//ff:func feature=sql type=command control=sequence
//ff:what TestRunSQL_JSONCov 테스트
package main

import (
	"path/filepath"
	"testing"
)

func TestRunSQL_JSONCov(t *testing.T) {
	dir := t.TempDir()
	outFile := filepath.Join(dir, "out.json")
	execSQL([]string{dir, "--json", "-o", outFile})
}
