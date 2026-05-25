//ff:func feature=hurl type=parse control=sequence
//ff:what TestParseEndpointID_NoSpace 테스트
package hurls

import "testing"

func TestParseEndpointID_NoSpace(t *testing.T) {
	method, path := parseEndpointID("GET")
	if method != "GET" || path != "" {
		t.Fatalf("got %q, %q", method, path)
	}
}
