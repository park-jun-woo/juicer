//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestTryParseRegisterBlueprint_NoPrefix 테스트
package flask

import "testing"

func TestTryParseRegisterBlueprint_NoPrefix(t *testing.T) {
	call, b := firstCall(t, `app.register_blueprint(api)`+"\n")
	v, p := tryParseRegisterBlueprint(call, b)
	if v != "api" || p != "" {
		t.Fatalf("got %q %q", v, p)
	}
}
