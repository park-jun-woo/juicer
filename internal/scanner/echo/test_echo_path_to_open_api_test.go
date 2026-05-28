//ff:func feature=scan type=test control=sequence
//ff:what TestEchoPathToOpenAPI_Param 테스트
package echo

import "testing"

func TestEchoPathToOpenAPI_Param(t *testing.T) {
	got := echoPathToOpenAPI("/api/users/:id")
	if got != "/api/users/{id}" {
		t.Fatalf("expected /api/users/{id}, got %s", got)
	}
}
