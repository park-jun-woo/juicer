//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestExtractURLParams_Direct 테스트
package flask

import "testing"

func TestExtractURLParams_Direct(t *testing.T) {
	params := extractURLParams("/users/<int:user_id>/posts/<post_id>/<path:rest>")
	if len(params) != 3 {
		t.Fatalf("expected 3 params, got %d: %v", len(params), params)
	}
	if params[0].name != "user_id" || params[0].converter != "int" {
		t.Errorf("param 0 = %+v", params[0])
	}
	if params[1].name != "post_id" || params[1].converter != "" {
		t.Errorf("param 1 = %+v", params[1])
	}
	if params[2].name != "rest" || params[2].converter != "path" {
		t.Errorf("param 2 = %+v", params[2])
	}
}
