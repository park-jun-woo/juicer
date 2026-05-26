//ff:func feature=scan type=test control=sequence
//ff:what TestRunScan_FrameworkGoginFlag --framework gogin 플래그 분기 테스트
package main

import "testing"

func TestRunScan_FrameworkGoginFlag(t *testing.T) {
	dir := setupMinimalGoProject(t)
	runScan([]string{"--framework", "gogin", dir})
}
