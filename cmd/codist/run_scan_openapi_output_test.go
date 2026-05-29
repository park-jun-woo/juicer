//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_OpenAPIOutput OpenAPI 출력 분기 테스트
package main

import "testing"

func TestRunScan_OpenAPIOutput(t *testing.T) {
	dir := setupMinimalGoProject(t)
	execScan([]string{"--openapi", dir})
}
