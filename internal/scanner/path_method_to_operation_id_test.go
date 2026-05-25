//ff:func feature=scan type=extract control=sequence
//ff:what TestPathMethodToOperationID_Basic 테스트
package scanner

import "testing"

func TestPathMethodToOperationID_Basic(t *testing.T) {
	got := pathMethodToOperationID("GET", "/api/v1/users")
	if got != "get_users" {
		t.Fatalf("expected get_users, got %s", got)
	}
}
