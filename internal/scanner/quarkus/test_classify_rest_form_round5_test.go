//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestClassifyRestForm_Round5 테스트
package quarkus

import "testing"

func TestClassifyRestForm_Round5(t *testing.T) {
	param, src := classifyFixture(t, `class C { void m(@RestForm String f) {} }`)
	var ep endpointInfo
	classifyRestForm(param, src, &ep, "String", "f")
	if len(ep.formParams) == 0 {
		t.Fatalf("formFields: %+v", ep.formParams)
	}
}
