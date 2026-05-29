//ff:func feature=scan type=command control=sequence
//ff:what TestRunDDL_EmptyDirNoError 테스트
package main

import (
	"path/filepath"
	"testing"
)

func TestRunDDL_EmptyDirNoError(t *testing.T) {
	// Empty dir with no SQL files — Parse returns empty map, no error
	dir := t.TempDir()
	execDDL([]string{"-o", filepath.Join(dir, "out"), dir})
}
