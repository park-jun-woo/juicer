//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestClassifyParam_RequestBody 테스트
package spring

import "testing"

func TestClassifyParam_RequestBody(t *testing.T) {
	p, src := firstParamS(t, `class C { void m(@RequestBody UserDto dto) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if ep.bodyType != "UserDto" {
		t.Fatalf("got %q", ep.bodyType)
	}
}
