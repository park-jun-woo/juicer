//ff:func feature=scan type=test control=sequence
//ff:what TestFiberPathToOpenAPI_Param 테스트
package fiber

import "testing"

func TestFiberPathToOpenAPI_Param(t *testing.T) {
	got := fiberPathToOpenAPI("/api/users/:id")
	if got != "/api/users/{id}" {
		t.Fatalf("expected /api/users/{id}, got %s", got)
	}
}
