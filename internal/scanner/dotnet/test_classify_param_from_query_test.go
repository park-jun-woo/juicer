//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestClassifyParam_FromQuery 테스트
package dotnet

import "testing"

func TestClassifyParam_FromQuery(t *testing.T) {
	p, src := firstParamCS(t, `class C { void m([FromQuery] string q) {} }`)
	ep := &endpointInfo{}
	classifyParam(p, src, ep)
	if len(ep.query) != 1 {
		t.Fatalf("got %+v", ep.query)
	}
}
