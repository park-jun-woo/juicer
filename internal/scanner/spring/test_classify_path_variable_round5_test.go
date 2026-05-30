//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestClassifyPathVariable_Round5 테스트
package spring

import "testing"

func TestClassifyPathVariable_Round5(t *testing.T) {
	param, src := sParam(t, `class C { void m(@PathVariable("id") Long id) {} }`)
	var ep endpointInfo
	classifyPathVariable(param, src, &ep, "Long", "id")
	if len(ep.params) != 1 || ep.params[0].Name != "id" {
		t.Fatalf("params: %+v", ep.params)
	}
}
