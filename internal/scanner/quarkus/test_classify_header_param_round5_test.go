//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestClassifyHeaderParam_Round5 테스트
package quarkus

import "testing"

func TestClassifyHeaderParam_Round5(t *testing.T) {
	param, src := classifyFixture(t, `class C { void m(@HeaderParam("X-Token") String token) {} }`)
	var ep endpointInfo
	classifyHeaderParam(param, src, &ep, "String", "token")
	if len(ep.headers) != 1 {
		t.Fatalf("headers: %+v", ep.headers)
	}
}
