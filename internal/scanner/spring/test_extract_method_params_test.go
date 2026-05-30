//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractMethodParams 테스트
package spring

import "testing"

func TestExtractMethodParams(t *testing.T) {
	fi := sFileInfo(t, `class C { void m(@PathVariable("id") Long id, @RequestParam("q") String q) {} }`)
	m := findAllByType(fi.root, "method_declaration")[0]
	ep := &endpointInfo{}
	extractMethodParams(m, fi.src, ep, nil, "", "")
	if len(ep.params) != 1 || len(ep.query) != 1 {
		t.Fatalf("got %+v %+v", ep.params, ep.query)
	}
}
