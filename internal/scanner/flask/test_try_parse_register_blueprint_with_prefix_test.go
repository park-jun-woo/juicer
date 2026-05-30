//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestTryParseRegisterBlueprint_WithPrefix 테스트
package flask

import "testing"

func TestTryParseRegisterBlueprint_WithPrefix(t *testing.T) {
	call, b := firstCall(t, `app.register_blueprint(users_bp, url_prefix="/v2")`+"\n")
	v, p := tryParseRegisterBlueprint(call, b)
	if v != "users_bp" || p != "/v2" {
		t.Fatalf("got %q %q", v, p)
	}
}
