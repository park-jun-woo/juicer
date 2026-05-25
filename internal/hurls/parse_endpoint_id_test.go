//ff:func feature=hurl type=parse control=sequence
//ff:what TestParseEndpointID 테스트
package hurls

import "testing"

func TestParseEndpointID(t *testing.T) {
	method, path := parseEndpointID("GET /api/health")
	if method != "GET" || path != "/api/health" {
		t.Fatalf("got %q, %q", method, path)
	}
}
