//ff:func feature=scan type=test control=selection topic=actix
//ff:what processServiceArg — web::scope/web::resource/기타 분기를 검증
package actix

import "testing"

func TestProcessServiceArg_Resource(t *testing.T) {
	src := []byte(`fn f() { web::resource("/users").route(web::get().to(list_users)); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".route")
	if call == nil {
		t.Fatal("no chain call")
	}
	var routes []builderRoute
	processServiceArg(call, src, "/api", &routes)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d: %+v", len(routes), routes)
	}
	if routes[0].path != "/api/users" {
		t.Errorf("path = %q, want /api/users", routes[0].path)
	}
	if routes[0].method != "GET" || routes[0].handler != "list_users" {
		t.Errorf("unexpected route: %+v", routes[0])
	}
}

func TestProcessServiceArg_Scope(t *testing.T) {
	src := []byte(`fn f() { web::scope("/v1").service(web::resource("/items").route(web::get().to(items))); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".service")
	if call == nil {
		t.Fatal("no scope chain call")
	}
	var routes []builderRoute
	processServiceArg(call, src, "/api", &routes)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d: %+v", len(routes), routes)
	}
	if routes[0].path != "/api/v1/items" {
		t.Errorf("path = %q, want /api/v1/items", routes[0].path)
	}
}

func TestProcessServiceArg_Other(t *testing.T) {
	// A call that is neither web::scope nor web::resource -> default: no routes.
	src := []byte(`fn f() { other::thing("/x"); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, "::thing")
	if call == nil {
		t.Fatal("no call")
	}
	var routes []builderRoute
	processServiceArg(call, src, "/api", &routes)
	if len(routes) != 0 {
		t.Fatalf("expected no routes, got %+v", routes)
	}
}
