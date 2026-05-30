//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestTryParseRegisterBlueprint_NoAttribute 테스트
package flask

import "testing"

func TestTryParseRegisterBlueprint_NoAttribute(t *testing.T) {

	call, b := firstCall(t, `register_blueprint(api)`+"\n")
	if v, _ := tryParseRegisterBlueprint(call, b); v != "" {
		t.Fatalf("plain call should be empty, got %q", v)
	}
}
