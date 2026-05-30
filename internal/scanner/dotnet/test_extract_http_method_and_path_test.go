//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractHTTPMethodAndPath 테스트
package dotnet

import "testing"

func TestExtractHTTPMethodAndPath(t *testing.T) {
	root, src := parseCS(t, `class C { [HttpGet("{id}")] public string Get(int id) { return ""; } }`)
	m := findAllByType(root, "method_declaration")[0]
	method, path, ok := extractHTTPMethodAndPath(m, src)
	if !ok || method != "GET" || path != "{id}" {
		t.Fatalf("got %q %q %v", method, path, ok)
	}
}
