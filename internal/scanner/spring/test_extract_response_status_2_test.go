//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractResponseStatus 테스트
package spring

import "testing"

func TestExtractResponseStatus(t *testing.T) {
	fi := sFileInfo(t, `class C { @ResponseStatus(HttpStatus.CREATED) void m() {} }`)
	m := findAllByType(fi.root, "method_declaration")[0]
	ep := &endpointInfo{}
	extractResponseStatus(m, fi.src, ep)
	if ep.statusCode != "201" {
		t.Fatalf("got %q", ep.statusCode)
	}
}
