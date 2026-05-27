//ff:func feature=ddl type=test control=sequence
//ff:what TestRunDDL_ParseErr 테스트
package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestRunDDL_ParseErr(t *testing.T) {
	if os.Getenv("RD_ERR_PARSE") == "1" {
		dir := t.TempDir()
		f := filepath.Join(dir, "001.up.sql")
		os.WriteFile(f, []byte("x"), 0o644)
		os.Chmod(f, 0o000)
		defer os.Chmod(f, 0o644)
		runDDL([]string{dir})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunDDL_ParseErr$")
	cmd.Env = append(os.Environ(), "RD_ERR_PARSE=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatal("expected non-zero exit")
}
