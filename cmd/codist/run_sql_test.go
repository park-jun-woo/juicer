//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_YAMLStdoutBranch YAML stdout 분기 테스트
package main

import "testing"

func TestRunSQL_YAMLStdoutBranch(t *testing.T) {
	dir := t.TempDir()
	runSQL([]string{dir})
}
