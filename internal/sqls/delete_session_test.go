//ff:func feature=sql type=parse control=sequence
//ff:what TestDeleteSession_NoSession 테스트
package sqls

import "testing"

func TestDeleteSession_NoSession(t *testing.T) {
	err := DeleteSession()
	if err != nil {
		t.Fatal(err)
	}
}
