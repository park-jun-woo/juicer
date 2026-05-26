//ff:func feature=sql type=test control=sequence
//ff:what TestCreateSession_EmptyRepoCov 테스트
package sqls

import "testing"

func TestCreateSession_EmptyRepoCov(t *testing.T) {
	err := createSession("", "/tmp/q")
	if err == nil {
		t.Fatal("expected error for empty repo")
	}
}
