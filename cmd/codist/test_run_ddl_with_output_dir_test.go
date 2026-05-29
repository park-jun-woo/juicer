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
	execDDL([]string{"-o", outDir, dir})
	matches, _ := filepath.Glob(filepath.Join(outDir, "*_orders.sql"))
	if len(matches) == 0 {
		t.Fatalf("expected a *_orders.sql file to exist in %s", outDir)
	}
}
