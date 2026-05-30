//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestClassifyParam_RestForm 테스트
package quarkus

import "testing"

func TestClassifyParam_RestForm(t *testing.T) {
	p, src := firstParam(t, `class R { void m(@RestForm("name") String name) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if len(ep.formParams) != 1 {
		t.Fatalf("got %+v", ep.formParams)
	}
}
