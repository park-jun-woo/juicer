//ff:func feature=scan type=test control=sequence
//ff:what TestEchoPathToOpenAPI_BareWildcard 테스트
package echo

import "testing"

func TestEchoPathToOpenAPI_BareWildcard(t *testing.T) {
	got := echoPathToOpenAPI("/swagger/*")
	if got != "/swagger/{wildcard}" {
		t.Fatalf("expected /swagger/{wildcard}, got %s", got)
	}
}
