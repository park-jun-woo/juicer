//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestClassifyQueryParam_Default 테스트
package quarkus

import "testing"

func TestClassifyQueryParam_Default(t *testing.T) {
	p, src := firstParam(t, `class R { void m(@DefaultValue("5") @QueryParam("limit") int limit) {} }`)
	ep := &endpointInfo{}
	classifyQueryParam(p, src, ep, "int", "limit")
	if len(ep.query) != 1 || ep.query[0].Default != "5" {
		t.Fatalf("got %+v", ep.query)
	}
}
