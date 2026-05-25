//ff:func feature=scan type=extract control=sequence
//ff:what TestPathMethodToOperationID_WithParam 테스트
package scanner

import "testing"

func TestPathMethodToOperationID_WithParam(t *testing.T) {
	got := pathMethodToOperationID("GET", "/api/v1/users/:id")
	if got != "get_users_id" {
		t.Fatalf("expected get_users_id, got %s", got)
	}
}
