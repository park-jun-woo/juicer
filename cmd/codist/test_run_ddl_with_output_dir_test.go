//ff:func feature=ddl type=test control=sequence
//ff:what TestRunDDL_WithOutputDir 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunDDL_WithOutputDir(t *testing.T) {
	dir := t.TempDir()
	sql := "CREATE TABLE orders (id INT PRIMARY KEY, total DECIMAL);\n"
	os.WriteFile(filepath.Join(dir, "001.up.sql"), []byte(sql), 0o644)
	outDir := filepath.Join(dir, "out")
	runDDL([]string{"-o", outDir, dir})
	if _, err := os.Stat(filepath.Join(outDir, "orders.sql")); err != nil {
		t.Fatalf("expected orders.sql to exist: %v", err)
	}
}
