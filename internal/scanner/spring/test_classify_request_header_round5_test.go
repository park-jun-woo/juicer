//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestClassifyRequestHeader_Round5 테스트
package spring

import "testing"

func TestClassifyRequestHeader_Round5(t *testing.T) {
	param, src := sParam(t, `class C { void m(@RequestHeader("X-Token") String token) {} }`)
	var ep endpointInfo
	classifyRequestHeader(param, src, &ep, "String", "token")
	if len(ep.headers) != 1 {
		t.Fatalf("headers: %+v", ep.headers)
	}
}
