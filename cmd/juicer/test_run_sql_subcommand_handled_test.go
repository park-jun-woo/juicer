//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_SubcommandHandled 테스트
package main

import "testing"

func TestRunSQL_SubcommandHandled(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	runSQL([]string{"status"})
}
