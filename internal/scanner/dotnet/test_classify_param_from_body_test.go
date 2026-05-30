//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestClassifyParam_FromBody 테스트
package dotnet

import "testing"

func TestClassifyParam_FromBody(t *testing.T) {
	p, src := firstParamCS(t, `class C { void m([FromBody] UserDto dto) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep)
	if ep.bodyType != "UserDto" {
		t.Fatalf("got %q", ep.bodyType)
	}
}
