//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_OpenAPIAll 테스트
package main

import "testing"

func TestRunScan_OpenAPIAll(t *testing.T) {
	dir := setupMinimalGoProject(t)
	runScan([]string{"-openapi", dir})
}
