//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestJoinPath 테스트
package dotnet

import "testing"

func TestJoinPath(t *testing.T) {
	if got := joinPath("/api", "users", "/{id}/"); got != "/api/users/{id}" {
		t.Fatalf("got %q", got)
	}
	if got := joinPath("", ""); got != "/" {
		t.Fatalf("got %q", got)
	}
}
