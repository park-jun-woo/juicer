//ff:func feature=sql type=test control=sequence
//ff:what TestCreateSession_EmptyQueriesCov 테스트
package sqls

import "testing"

func TestCreateSession_EmptyQueriesCov(t *testing.T) {
	err := createSession("/tmp/r", "")
	if err == nil {
		t.Fatal("expected error for empty queries")
	}
}
