//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractHTTPMethodAndPath_None 테스트
package spring

import "testing"

func TestExtractHTTPMethodAndPath_None(t *testing.T) {
	m, src := firstMethodS(t, `class C { public String helper() { return ""; } }`)
	if _, _, ok := extractHTTPMethodAndPath(m, src); ok {
		t.Fatal("expected false")
	}
}
