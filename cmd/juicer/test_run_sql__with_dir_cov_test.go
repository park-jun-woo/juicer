//ff:func feature=sql type=command control=sequence
//ff:what TestRunSQL_WithDirCov 테스트
package main

import "testing"

func TestRunSQL_WithDirCov(t *testing.T) {
	dir := t.TempDir()
	runSQL([]string{dir})
}
