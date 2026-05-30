//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestClassifyAnnotatedDepends_AliasMatch 테스트
package fastapi

import "testing"

func TestClassifyAnnotatedDepends_AliasMatch(t *testing.T) {
	ri := &routeInfo{}
	classifyAnnotatedDepends("dep", "AuthDep", map[string]string{"AuthDep": "verify_token"}, ri)
	if len(ri.middleware) != 1 || ri.middleware[0] != "verify_token" {
		t.Fatalf("got %+v", ri.middleware)
	}
}
