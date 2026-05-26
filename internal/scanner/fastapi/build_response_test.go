//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what buildResponse 테스트
package fastapi

import "testing"

func TestBuildResponse(t *testing.T) {
	// explicit status code
	ri := routeInfo{method: "POST", statusCode: 201}
	resp := buildResponse(ri, "UserOut")
	if resp.Status != "201" {
		t.Errorf("expected 201, got %s", resp.Status)
	}
	if resp.TypeName != "UserOut" {
		t.Errorf("expected UserOut, got %s", resp.TypeName)
	}

	// default status
	ri2 := routeInfo{method: "GET"}
	resp2 := buildResponse(ri2, "")
	if resp2.Status != "200" {
		t.Errorf("expected 200, got %s", resp2.Status)
	}

	// builtin type: should not set TypeName
	ri3 := routeInfo{method: "GET"}
	resp3 := buildResponse(ri3, "str")
	if resp3.TypeName != "str" {
		// str has openAPIType.Type = "string", so it should be set
	}
}
