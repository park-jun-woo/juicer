//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestJoinPath 테스트
package quarkus

import "testing"

func TestJoinPath(t *testing.T) {
	if got := joinPath("/api", "users", "/{id}/"); got != "/api/users/{id}" {
		t.Fatalf("got %q", got)
	}
	if got := joinPath("", "", ""); got != "/" {
		t.Fatalf("got %q", got)
	}
	if got := joinPath("/users"); got != "/users" {
		t.Fatalf("got %q", got)
	}
}
