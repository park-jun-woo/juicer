//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestClassifyParam_Header 테스트
package quarkus

import "testing"

func TestClassifyParam_Header(t *testing.T) {
	p, src := firstParam(t, `class R { void m(@HeaderParam("X-Token") String tok) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if len(ep.headers) != 1 {
		t.Fatalf("got %+v", ep.headers)
	}
}
