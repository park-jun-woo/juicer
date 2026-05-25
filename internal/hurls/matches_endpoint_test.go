//ff:func feature=hurl type=parse control=sequence
//ff:what TestMatchesEndpoint_Host 테스트
package hurls

import "testing"

func TestMatchesEndpoint_Host(t *testing.T) {
	content := "GET {{host}}/api/health\nHTTP 200\n"
	if !matchesEndpoint(content, "GET", "/api/health") {
		t.Fatal("expected true")
	}
}
