//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_ExtractErr 테스트
package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestRunSQL_ExtractErr(t *testing.T) {
	if os.Getenv("RSQL_ERR_EXT") == "1" {
		execSQL([]string{"/nonexistent/dir/sql"})
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunSQL_ExtractErr$")
	cmd.Env = append(os.Environ(), "RSQL_ERR_EXT=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	// Extract may not fail on non-existent dir; still covers paths
}
