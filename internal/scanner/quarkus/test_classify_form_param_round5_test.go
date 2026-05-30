//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestClassifyFormParam_Round5 테스트
package quarkus

import "testing"

func TestClassifyFormParam_Round5(t *testing.T) {
	param, src := classifyFixture(t, `class C { void m(@FormParam("f") String f) {} }`)
	var ep endpointInfo
	classifyFormParam(param, src, &ep, "String", "f")
	if len(ep.formParams) == 0 {
		t.Fatalf("formFields: %+v", ep.formParams)
	}
}
