//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_YAMLDefault 테스트
package main

import "testing"

func TestRunScan_YAMLDefault(t *testing.T) {
	dir := setupMinimalGoProject(t)
	runScan([]string{dir})
}
