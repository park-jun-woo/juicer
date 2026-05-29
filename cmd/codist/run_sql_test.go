//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_YAMLStdoutBranch YAML stdout 분기 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunSQL_YAMLStdoutBranch(t *testing.T) {
	// YAML to stdout (skeleton mode, dir positional arg)
	dir := t.TempDir()
	execSQL([]string{dir})

	// JSON to stdout
	execSQL([]string{"--json", dir})

	// output-file branch
	out := filepath.Join(t.TempDir(), "queries.yaml")
	execSQL([]string{"-o", out, dir})
	if _, err := os.Stat(out); err != nil {
		t.Fatalf("expected output file written: %v", err)
	}

	// subcommand-handled branch (status dispatches to handleSQLSubcommand)
	// Remaining uncovered lines (32, 43, 49) are os.Exit error paths that
	// cannot be triggered in-process without aborting the test binary.
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	execSQL([]string{"status"})
}
