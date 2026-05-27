//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_DirArgJSON 테스트
package main

import "testing"

func TestRunSQL_DirArgJSON(t *testing.T) {
	dir := t.TempDir()
	runSQL([]string{"-json", dir})
}
