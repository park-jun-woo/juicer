//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestParseRouteArg_NoRoute 테스트
package actix

import "testing"

func TestParseRouteArg_NoRoute(t *testing.T) {

	src := []byte(`fn f() { x.route(plain_arg); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".route")
	args := findChildByType(call, "arguments")
	method, handler := parseRouteArg(args, src)
	if method != "" || handler != "" {
		t.Fatalf("expected empty, got (%q, %q)", method, handler)
	}
}
