//ff:func feature=scan type=extract control=sequence
//ff:what TestGinPathToOpenAPI_NoParam 테스트
package scanner

import "testing"

func TestGinPathToOpenAPI_NoParam(t *testing.T) {
	got := ginPathToOpenAPI("/api/users")
	if got != "/api/users" {
		t.Fatalf("expected /api/users, got %s", got)
	}
}
