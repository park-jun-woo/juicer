//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestEchoPathToOpenAPI 테스트
package echo

import "testing"

func TestEchoPathToOpenAPI(t *testing.T) {
	if got := echoPathToOpenAPI("/users/:id"); got != "/users/{id}" {
		t.Fatalf("got %q", got)
	}
	if got := echoPathToOpenAPI("/files/*"); got != "/files/{wildcard}" {
		t.Fatalf("got %q", got)
	}
	if got := echoPathToOpenAPI("/files/*path"); got != "/files/{path}" {
		t.Fatalf("got %q", got)
	}
}
