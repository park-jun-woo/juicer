//ff:func feature=ddl type=test control=sequence
//ff:what TestRunDDL_ParseError 테스트
package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestRunDDL_ParseError(t *testing.T) {
	if os.Getenv("TEST_SUBPROCESS_DDL_PARSE") == "1" {
		// Create a .up.sql file that is unreadable to trigger Parse read error
		dir := t.TempDir()
		f := filepath.Join(dir, "001.up.sql")
		os.WriteFile(f, []byte("x"), 0o644)
		os.Chmod(f, 0o000)
		defer os.Chmod(f, 0o644)
		runDDL([]string{dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunDDL_ParseError$")
	cmd.Env = append(os.Environ(), "TEST_SUBPROCESS_DDL_PARSE=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit from subprocess")
}
