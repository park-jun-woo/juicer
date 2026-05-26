//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_JSONOutput JSON 출력 분기 테스트
package main

import "testing"

func TestRunSQL_JSONOutput(t *testing.T) {
	dir := t.TempDir()
	runSQL([]string{"--json", dir})
}
