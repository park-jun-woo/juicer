//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_JSONOutputAll 테스트
package main

import "testing"

func TestRunScan_JSONOutputAll(t *testing.T) {
	dir := setupMinimalGoProject(t)
	execScan([]string{"--json", dir})
}
