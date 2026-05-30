//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestClassifyParam_Query 테스트
package quarkus

import "testing"

func TestClassifyParam_Query(t *testing.T) {
	p, src := firstParam(t, `class R { void m(@QueryParam("limit") int limit) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if len(ep.query) != 1 || ep.query[0].Name != "limit" {
		t.Fatalf("got %+v", ep.query)
	}
}
