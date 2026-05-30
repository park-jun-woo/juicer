//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractHTTPMethodAndPath_GetMapping 테스트
package spring

import "testing"

func TestExtractHTTPMethodAndPath_GetMapping(t *testing.T) {
	m, src := firstMethodS(t, `class C { @GetMapping("/users") public String list() { return ""; } }`)
	method, path, ok := extractHTTPMethodAndPath(m, src)
	if !ok || method != "GET" || path != "/users" {
		t.Fatalf("got %q %q %v", method, path, ok)
	}
}
