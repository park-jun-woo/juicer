//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestClassifyParam_FromRoute 테스트
package dotnet

import "testing"

func TestClassifyParam_FromRoute(t *testing.T) {
	p, src := firstParamCS(t, `class C { void m([FromRoute] int id) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep)
	if len(ep.params) != 1 {
		t.Fatalf("got %+v", ep.params)
	}
}
