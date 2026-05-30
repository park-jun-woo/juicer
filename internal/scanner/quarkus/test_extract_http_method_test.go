//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractHTTPMethod 테스트
package quarkus

import "testing"

func TestExtractHTTPMethod(t *testing.T) {
	root, src := parseQ(t, `class R { @GET public String list() { return ""; } }`)
	m := findAllByType(root, "method_declaration")[0]
	method, ok := extractHTTPMethod(m, src)
	if !ok || method != "GET" {
		t.Fatalf("got %q %v", method, ok)
	}
}
