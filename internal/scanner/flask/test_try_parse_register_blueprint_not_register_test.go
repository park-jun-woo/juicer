//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestTryParseRegisterBlueprint_NotRegister 테스트
package flask

import "testing"

func TestTryParseRegisterBlueprint_NotRegister(t *testing.T) {
	call, b := firstCall(t, `app.route("/x")`+"\n")
	if v, p := tryParseRegisterBlueprint(call, b); v != "" || p != "" {
		t.Fatalf("non-register should be empty, got %q %q", v, p)
	}
}
