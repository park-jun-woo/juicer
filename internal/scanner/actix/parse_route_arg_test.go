//ff:func feature=scan type=test control=sequence topic=actix
//ff:what parseRouteArg — .route() 인자에서 (method, handler) 파싱 분기를 검증
package actix

import "testing"

func TestParseRouteArg(t *testing.T) {
	src := []byte(`fn f() { x.route(web::get().to(get_user)); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".route")
	if call == nil {
		t.Fatal("no .route call")
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		t.Fatal("no arguments")
	}
	method, handler := parseRouteArg(args, src)
	if method != "GET" {
		t.Errorf("method = %q, want GET", method)
	}
	if handler != "get_user" {
		t.Errorf("handler = %q, want get_user", handler)
	}
}

func TestParseRouteArg_NoRoute(t *testing.T) {
	// arguments that are not a web::<method>().to(...) chain -> "", "".
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
