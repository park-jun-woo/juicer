//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestClassifyParam_Path 테스트
package quarkus

import "testing"

func TestClassifyParam_Path(t *testing.T) {
	p, src := firstParam(t, `class R { void m(@PathParam("id") String id) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if len(ep.params) != 1 || ep.params[0].Name != "id" {
		t.Fatalf("got %+v", ep.params)
	}
}
