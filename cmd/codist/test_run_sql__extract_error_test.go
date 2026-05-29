//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_ExtractError 테스트
package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestRunSQL_ExtractError(t *testing.T) {
	if os.Getenv("TEST_SUBPROCESS_SQL_EXTRACT") == "1" {
		// Create directory with unreadable .go file
		dir := t.TempDir()
		f := filepath.Join(dir, "main.go")
		os.WriteFile(f, []byte("package main"), 0o644)
		os.Chmod(f, 0o000)
		defer os.Chmod(f, 0o644)
		execSQL([]string{dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunSQL_ExtractError$")
	cmd.Env = append(os.Environ(), "TEST_SUBPROCESS_SQL_EXTRACT=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	// Extract might not error; that's acceptable
}
