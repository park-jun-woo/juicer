//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_SubcommandBranch 테스트
package main

import "testing"

func TestRunSQL_SubcommandBranch(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	execSQL([]string{"status"})
}
