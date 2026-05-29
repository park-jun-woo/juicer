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
	execSQL([]string{"--json", dir})

	// Output to file branch
	out := filepath.Join(t.TempDir(), "out.yaml")
	execSQL([]string{"-o", out, dir})

	// Subcommand dispatch (status via setupSQLSession)
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	execSQL([]string{"status"})

	// Default dir branch (no args beyond flags)
	execSQL([]string{})
}
