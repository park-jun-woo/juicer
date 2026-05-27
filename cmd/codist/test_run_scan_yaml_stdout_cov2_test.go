//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_YAMLStdoutCov2 테스트
package main

import "testing"

func TestRunScan_YAMLStdoutCov2(t *testing.T) {
	dir := setupMinimalGoProject(t)
	runScan([]string{dir})
}
