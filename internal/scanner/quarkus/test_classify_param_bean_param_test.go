//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestClassifyParam_BeanParam 테스트
package quarkus

import "testing"

func TestClassifyParam_BeanParam(t *testing.T) {
	p, src := firstParam(t, `class R { void m(@BeanParam Filters f) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if ep.formType != "Filters" {
		t.Fatalf("formType: %q", ep.formType)
	}
}
