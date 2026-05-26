//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_AllBranches 테스트
package main

import (
	"path/filepath"
	"testing"
)

func TestRunSQL_AllBranches(t *testing.T) {
	dir := t.TempDir()

	// JSON output branch
	runSQL([]string{"-json", dir})

	// Output to file branch
	out := filepath.Join(t.TempDir(), "out.yaml")
	runSQL([]string{"-o", out, dir})

	// Subcommand dispatch (status via setupSQLSession)
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	runSQL([]string{"status"})

	// Default dir branch (no args beyond flags)
	runSQL([]string{})
}
