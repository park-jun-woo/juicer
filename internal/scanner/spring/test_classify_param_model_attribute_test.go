//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestClassifyParam_ModelAttribute 테스트
package spring

import "testing"

func TestClassifyParam_ModelAttribute(t *testing.T) {
	p, src := firstParamS(t, `class C { void m(@ModelAttribute Filters f) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if ep.formType != "Filters" {
		t.Fatalf("got %q", ep.formType)
	}
}
