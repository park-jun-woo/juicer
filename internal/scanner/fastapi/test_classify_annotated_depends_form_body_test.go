//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestClassifyAnnotatedDepends_FormBody 테스트
package fastapi

import "testing"

func TestClassifyAnnotatedDepends_FormBody(t *testing.T) {
	ri := &routeInfo{}
	classifyAnnotatedDepends("form", "Annotated[OAuth2PasswordRequestForm, Depends()]", nil, ri)
	if ri.bodyType != "OAuth2PasswordRequestForm" || ri.bodyVarName != "form" {
		t.Fatalf("got %+v", ri)
	}
}
