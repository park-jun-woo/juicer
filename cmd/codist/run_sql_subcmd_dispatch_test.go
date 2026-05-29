//ff:func feature=sql type=test control=sequence
//ff:what TestRunSQL_SubcmdDispatch 서브커맨드 디스패치 분기 테스트
package main

import "testing"

func TestRunSQL_SubcmdDispatch(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	execSQL([]string{"status"})
}
