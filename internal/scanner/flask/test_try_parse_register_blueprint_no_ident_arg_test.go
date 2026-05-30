//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestTryParseRegisterBlueprint_NoIdentArg 테스트
package flask

import "testing"

func TestTryParseRegisterBlueprint_NoIdentArg(t *testing.T) {

	call, b := firstCall(t, `app.register_blueprint("not_an_ident")`+"\n")
	if v, _ := tryParseRegisterBlueprint(call, b); v != "" {
		t.Fatalf("string arg should yield empty var, got %q", v)
	}
}
