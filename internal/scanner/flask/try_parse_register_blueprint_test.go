//ff:func feature=scan type=test control=sequence topic=flask
//ff:what tryParseRegisterBlueprint 테스트
package flask

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstCall(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	calls := findAllByType(root, "call")
	if len(calls) == 0 {
		t.Fatal("no call")
	}
	return calls[0], b
}

func TestTryParseRegisterBlueprint_WithPrefix(t *testing.T) {
	call, b := firstCall(t, `app.register_blueprint(users_bp, url_prefix="/v2")`+"\n")
	v, p := tryParseRegisterBlueprint(call, b)
	if v != "users_bp" || p != "/v2" {
		t.Fatalf("got %q %q", v, p)
	}
}

func TestTryParseRegisterBlueprint_NoPrefix(t *testing.T) {
	call, b := firstCall(t, `app.register_blueprint(api)`+"\n")
	v, p := tryParseRegisterBlueprint(call, b)
	if v != "api" || p != "" {
		t.Fatalf("got %q %q", v, p)
	}
}

func TestTryParseRegisterBlueprint_NotRegister(t *testing.T) {
	call, b := firstCall(t, `app.route("/x")`+"\n")
	if v, p := tryParseRegisterBlueprint(call, b); v != "" || p != "" {
		t.Fatalf("non-register should be empty, got %q %q", v, p)
	}
}

func TestTryParseRegisterBlueprint_NoAttribute(t *testing.T) {
	// plain function call, no attribute -> empty
	call, b := firstCall(t, `register_blueprint(api)`+"\n")
	if v, _ := tryParseRegisterBlueprint(call, b); v != "" {
		t.Fatalf("plain call should be empty, got %q", v)
	}
}

func TestTryParseRegisterBlueprint_NestedAttribute(t *testing.T) {
	// app.bp.register_blueprint(...) still ends with .register_blueprint
	call, b := firstCall(t, `app.bp.register_blueprint(sub)`+"\n")
	if v, _ := tryParseRegisterBlueprint(call, b); v != "sub" {
		t.Fatalf("nested attr register: got %q", v)
	}
}

func TestTryParseRegisterBlueprint_NoIdentArg(t *testing.T) {
	// register_blueprint with a string arg (not an identifier) -> empty var
	call, b := firstCall(t, `app.register_blueprint("not_an_ident")`+"\n")
	if v, _ := tryParseRegisterBlueprint(call, b); v != "" {
		t.Fatalf("string arg should yield empty var, got %q", v)
	}
}
