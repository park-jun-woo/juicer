//ff:func feature=ratchet type=test control=sequence
//ff:what TestRunSQLNext_Happy 정상 실행 테스트
package main

import "testing"

func TestRunSQLNext_Happy(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	runSQLNext([]string{})
}
