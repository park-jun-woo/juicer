//ff:func feature=scan type=test control=sequence
//ff:what TestFiberPathToOpenAPI_NoParam 테스트
package fiber

import "testing"

func TestFiberPathToOpenAPI_NoParam(t *testing.T) {
	got := fiberPathToOpenAPI("/api/users")
	if got != "/api/users" {
		t.Fatalf("expected /api/users, got %s", got)
	}
}
