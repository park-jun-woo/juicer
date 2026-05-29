//ff:func feature=ratchet type=test control=sequence
//ff:what TestRunSQLNext_SessionExists 테스트
package main

import "testing"

func TestRunSQLNext_SessionExists(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	execSQLNext([]string{})
}
