//ff:func feature=hurl type=parse control=sequence
//ff:what TestMatchesEndpoint_HTTPS 테스트
package hurls

import "testing"

func TestMatchesEndpoint_HTTPS(t *testing.T) {
	content := "GET https://{{host}}/api/health\nHTTP 200\n"
	if !matchesEndpoint(content, "GET", "/api/health") {
		t.Fatal("expected true")
	}
}
