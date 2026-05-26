//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_JSONBranch 테스트
package main

import "testing"

func TestRunScan_JSONBranch(t *testing.T) {
	dir := setupMinimalGoProject(t)
	runScan([]string{"-json", dir})
}
