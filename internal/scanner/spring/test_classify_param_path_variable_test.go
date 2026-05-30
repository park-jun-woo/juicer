//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestClassifyParam_PathVariable 테스트
package spring

import "testing"

func TestClassifyParam_PathVariable(t *testing.T) {
	p, src := firstParamS(t, `class C { void m(@PathVariable("id") Long id) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if len(ep.params) != 1 || ep.params[0].Name != "id" {
		t.Fatalf("got %+v", ep.params)
	}
}
