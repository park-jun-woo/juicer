//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestClassifyPathParam_Round5 테스트
package quarkus

import "testing"

func TestClassifyPathParam_Round5(t *testing.T) {
	param, src := classifyFixture(t, `class C { void m(@PathParam("id") String id) {} }`)
	var ep endpointInfo
	classifyPathParam(param, src, &ep, "String", "id")
	if len(ep.params) != 1 || ep.params[0].Name != "id" {
		t.Fatalf("params: %+v", ep.params)
	}
}
