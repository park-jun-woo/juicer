//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestClassifyParam_RequestHeader 테스트
package spring

import "testing"

func TestClassifyParam_RequestHeader(t *testing.T) {
	p, src := firstParamS(t, `class C { void m(@RequestHeader("X-Token") String tok) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if len(ep.headers) != 1 {
		t.Fatalf("got %+v", ep.headers)
	}
}
