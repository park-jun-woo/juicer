//ff:func feature=sql type=session control=sequence
//ff:what 테스트 헬퍼 — 테스트 세션 저장
package sqls

import "testing"

func setupTestSession(t *testing.T, sess *Session) {
	t.Helper()
	if err := SaveSession(sess); err != nil {
		t.Fatal(err)
	}
}
