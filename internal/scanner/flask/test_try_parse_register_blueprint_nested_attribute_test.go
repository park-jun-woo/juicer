//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestTryParseRegisterBlueprint_NestedAttribute 테스트
package flask

import "testing"

func TestTryParseRegisterBlueprint_NestedAttribute(t *testing.T) {

	call, b := firstCall(t, `app.bp.register_blueprint(sub)`+"\n")
	if v, _ := tryParseRegisterBlueprint(call, b); v != "sub" {
		t.Fatalf("nested attr register: got %q", v)
	}
}
