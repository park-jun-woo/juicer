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
	runScan([]string{"-json", dir})

	// OpenAPI output branch
	runScan([]string{"-openapi", dir})

	// Output to file branch
	out := filepath.Join(t.TempDir(), "out.yaml")
	runScan([]string{"-o", out, dir})

	// NestJS framework branch
	nestDir := t.TempDir()
	runScan([]string{"-framework", "nestjs", nestDir})

	// Explicit framework flag
	runScan([]string{"-framework", "gogin", dir})
}
