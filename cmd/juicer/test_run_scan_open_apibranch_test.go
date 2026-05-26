//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_OpenAPIBranch 테스트
package main

import (
	"path/filepath"
	"testing"
)

func TestRunScan_OpenAPIBranch(t *testing.T) {
	dir := setupMinimalGoProject(t)
	outFile := filepath.Join(dir, "openapi.yaml")
	runScan([]string{"-openapi", "-o", outFile, dir})
}
