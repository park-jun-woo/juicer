//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_SubcommandNext 테스트
package main

import "testing"

func TestRunSQL_SubcommandNext(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	execSQL([]string{"next"})
}
