//ff:func feature=hurl type=parse control=sequence
//ff:what TestMatchesEndpoint_NoMatch 테스트
package hurls

import "testing"

func TestMatchesEndpoint_NoMatch(t *testing.T) {
	content := "GET {{host}}/api/users\nHTTP 200\n"
	if matchesEndpoint(content, "GET", "/api/health") {
		t.Fatal("expected false")
	}
}
