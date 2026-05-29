//ff:func feature=scan type=command control=sequence
//ff:what TestRunDDL_Stdout 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunDDL_Stdout(t *testing.T) {
	dir := t.TempDir()
	sql := `CREATE TABLE items (id INT);`
	os.WriteFile(filepath.Join(dir, "001_init.up.sql"), []byte(sql), 0o644)
	execDDL([]string{dir})
}
