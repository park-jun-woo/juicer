//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractHTTPMethodAndPath_None 테스트
package dotnet

import "testing"

func TestExtractHTTPMethodAndPath_None(t *testing.T) {
	root, src := parseCS(t, `class C { public string Helper() { return ""; } }`)
	m := findAllByType(root, "method_declaration")[0]
	if _, _, ok := extractHTTPMethodAndPath(m, src); ok {
		t.Fatal("expected false")
	}
}
