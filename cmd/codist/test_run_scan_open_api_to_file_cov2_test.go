//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_OpenAPIToFileCov2 테스트
package main

import (
	"path/filepath"
	"testing"
)

func TestRunScan_OpenAPIToFileCov2(t *testing.T) {
	dir := setupMinimalGoProject(t)
	outFile := filepath.Join(dir, "out.yaml")
	execScan([]string{"--openapi", "-o", outFile, dir})
}
