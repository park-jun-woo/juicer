//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestClassifyParam_RequestParam 테스트
package spring

import "testing"

func TestClassifyParam_RequestParam(t *testing.T) {
	p, src := firstParamS(t, `class C { void m(@RequestParam("q") String q) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if len(ep.query) != 1 {
		t.Fatalf("got %+v", ep.query)
	}
}
