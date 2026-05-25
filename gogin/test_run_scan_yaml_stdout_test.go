//ff:func feature=scan type=command control=sequence
//ff:what TestRunScan_YAML_Stdout 테스트
package main

import (
	"testing"
)

func TestRunScan_YAML_Stdout(t *testing.T) {
	dir := setupMinimalGoProject(t)
	// Without -o flag, writes to stdout
	runScan([]string{dir})
}
