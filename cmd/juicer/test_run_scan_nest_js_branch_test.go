//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_NestJSBranch 테스트
package main

import "testing"

func TestRunScan_NestJSBranch(t *testing.T) {
	dir := t.TempDir()
	runScan([]string{"-framework", "nestjs", dir})
}
