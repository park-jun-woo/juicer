//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_FrameworkFlag 프레임워크 플래그 테스트
package main

import "testing"

func TestRunScan_FrameworkFlag(t *testing.T) {
	dir := setupMinimalGoProject(t)
	execScan([]string{"--framework", "gogin", dir})
}
