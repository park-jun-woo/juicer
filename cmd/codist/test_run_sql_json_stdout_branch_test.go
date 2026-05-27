//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_JSONStdoutBranch 테스트
package main

import "testing"

func TestRunSQL_JSONStdoutBranch(t *testing.T) {
	dir := t.TempDir()
	runSQL([]string{"-json", dir})
}
