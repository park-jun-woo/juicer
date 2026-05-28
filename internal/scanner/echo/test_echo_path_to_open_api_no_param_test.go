//ff:func feature=scan type=test control=sequence
//ff:what TestEchoPathToOpenAPI_NoParam 테스트
package echo

import "testing"

func TestEchoPathToOpenAPI_NoParam(t *testing.T) {
	got := echoPathToOpenAPI("/api/users")
	if got != "/api/users" {
		t.Fatalf("expected /api/users, got %s", got)
	}
}
