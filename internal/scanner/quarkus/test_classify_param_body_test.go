//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestClassifyParam_Body 테스트
package quarkus

import "testing"

func TestClassifyParam_Body(t *testing.T) {
	p, src := firstParam(t, `class R { void m(UserDto dto) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep, nil, "", "")
	if ep.bodyType != "UserDto" {
		t.Fatalf("bodyType: %q", ep.bodyType)
	}
}
