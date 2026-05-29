//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_AllBranches 테스트
package main

import (
	"path/filepath"
	"testing"
)

func TestRunScan_AllBranches(t *testing.T) {
	dir := setupMinimalGoProject(t)

	// JSON output branch
	execScan([]string{"--json", dir})

	// OpenAPI output branch
	execScan([]string{"--openapi", dir})

	// Output to file branch
	out := filepath.Join(t.TempDir(), "out.yaml")
	execScan([]string{"-o", out, dir})

	// NestJS framework branch
	nestDir := t.TempDir()
	execScan([]string{"--framework", "nestjs", nestDir})

	// Explicit framework flag
	execScan([]string{"--framework", "gogin", dir})
}
