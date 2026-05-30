//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what classifyAnnotatedDepends: alias미들웨어 / Depends없음 / form body / 일반 미들웨어
package fastapi

import "testing"

func TestClassifyAnnotatedDepends_AliasMatch(t *testing.T) {
	ri := &routeInfo{}
	classifyAnnotatedDepends("dep", "AuthDep", map[string]string{"AuthDep": "verify_token"}, ri)
	if len(ri.middleware) != 1 || ri.middleware[0] != "verify_token" {
		t.Fatalf("got %+v", ri.middleware)
	}
}

func TestClassifyAnnotatedDepends_NoDepends(t *testing.T) {
	ri := &routeInfo{}
	classifyAnnotatedDepends("x", "Annotated[int, Query()]", nil, ri)
	if len(ri.middleware) != 0 || ri.bodyType != "" {
		t.Fatalf("expected no-op, got %+v", ri)
	}
}

func TestClassifyAnnotatedDepends_FormBody(t *testing.T) {
	ri := &routeInfo{}
	classifyAnnotatedDepends("form", "Annotated[OAuth2PasswordRequestForm, Depends()]", nil, ri)
	if ri.bodyType != "OAuth2PasswordRequestForm" || ri.bodyVarName != "form" {
		t.Fatalf("got %+v", ri)
	}
}

func TestClassifyAnnotatedDepends_Middleware(t *testing.T) {
	ri := &routeInfo{}
	classifyAnnotatedDepends("dep", "Annotated[User, Depends(get_current_user)]", nil, ri)
	if len(ri.middleware) != 1 || ri.middleware[0] != "get_current_user" {
		t.Fatalf("got %+v", ri.middleware)
	}
}
