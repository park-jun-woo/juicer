package sqls

import "testing"

func TestDeleteSession_NoSession(t *testing.T) {
	err := DeleteSession()
	if err != nil {
		t.Fatal(err)
	}
}
