//ff:func feature=scan type=test control=sequence topic=actix
//ff:what processServiceCallArgs — arguments 내 call_expression 인자만 처리함을 검증
package actix

import "testing"

func TestProcessServiceCallArgs(t *testing.T) {
	src := []byte(`fn f() { cfg.service(web::resource("/users").route(web::get().to(list))); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".service")
	if call == nil {
		t.Fatal("no .service call")
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		t.Fatal("no arguments")
	}
	var routes []builderRoute
	processServiceCallArgs(args, src, "", &routes)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d: %+v", len(routes), routes)
	}
	if routes[0].path != "/users" || routes[0].handler != "list" {
		t.Errorf("unexpected route: %+v", routes[0])
	}
}

func TestProcessServiceCallArgs_NoCall(t *testing.T) {
	// arguments whose child is a plain identifier (not call_expression).
	src := []byte(`fn f() { cfg.service(handler); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".service")
	args := findChildByType(call, "arguments")
	var routes []builderRoute
	processServiceCallArgs(args, src, "", &routes)
	if len(routes) != 0 {
		t.Fatalf("expected no routes, got %+v", routes)
	}
}
