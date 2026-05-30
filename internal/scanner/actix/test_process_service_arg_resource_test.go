//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestProcessServiceArg_Resource 테스트
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
