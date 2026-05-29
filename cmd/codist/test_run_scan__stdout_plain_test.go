//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_StdoutPlain 테스트
package main

import "testing"

func TestRunScan_StdoutPlain(t *testing.T) {
	dir := setupMinimalGoProject(t)
	execScan([]string{dir})
}
