//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractMethodParams 테스트
package quarkus

import "testing"

func TestExtractMethodParams(t *testing.T) {
	fi := qFileInfo(t, `class R { void m(@PathParam("id") Long id, @QueryParam("q") String q) {} }`)
	m := findAllByType(fi.root, "method_declaration")[0]
	ep := &endpointInfo{}
	extractMethodParams(m, fi.src, ep, nil, "", "")
	if len(ep.params) != 1 || len(ep.query) != 1 {
		t.Fatalf("got params=%+v query=%+v", ep.params, ep.query)
	}
}
