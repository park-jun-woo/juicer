//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_DispatchSubcmd 서브커맨드 디스패치 테스트
package main

import "testing"

func TestRunSQL_DispatchSubcmd(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	execSQL([]string{"status"})
}
