//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestClassifyRequestParam_Round5 테스트
package spring

import "testing"

func TestClassifyRequestParam_Round5(t *testing.T) {
	param, src := sParam(t, `class C { void m(@RequestParam("q") String q) {} }`)
	var ep endpointInfo
	classifyRequestParam(param, src, &ep, "String", "q", map[string]string{}, "C.java", "/abs")
	if len(ep.query) != 1 || ep.query[0].Name != "q" {
		t.Fatalf("query: %+v", ep.query)
	}
}
