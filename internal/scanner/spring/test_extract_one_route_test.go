//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractOneRoute 테스트
package spring

import "testing"

func TestExtractOneRoute(t *testing.T) {
	fi := sFileInfo(t, sampleController)
	m := findAllByType(fi.root, "method_declaration")[0]
	ep, ok := extractOneRoute(m, fi)
	if !ok || ep.method != "GET" || ep.path != "/{id}" || ep.handler != "get" {
		t.Fatalf("got %+v ok=%v", ep, ok)
	}
	if len(ep.params) != 1 {
		t.Fatalf("params: %+v", ep.params)
	}
}
