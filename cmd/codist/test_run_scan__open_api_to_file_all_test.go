//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_OpenAPIToFileAll 테스트
package main

import (
	"path/filepath"
	"testing"
)

func TestRunScan_OpenAPIToFileAll(t *testing.T) {
	dir := setupMinimalGoProject(t)
	outFile := filepath.Join(dir, "api.yaml")
	execScan([]string{"--openapi", "-o", outFile, dir})
}
