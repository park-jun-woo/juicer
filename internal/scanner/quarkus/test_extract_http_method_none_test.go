//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractHTTPMethod_None 테스트
package quarkus

import "testing"

func TestExtractHTTPMethod_None(t *testing.T) {
	root, src := parseQ(t, `class R { public String list() { return ""; } }`)
	m := findAllByType(root, "method_declaration")[0]
	if _, ok := extractHTTPMethod(m, src); ok {
		t.Fatal("expected false")
	}
}
