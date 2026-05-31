//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestProcessServiceArg_Scope 테스트
package actix

import "testing"

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
	processServiceArg(call, src, "/api", &routes, nil, map[string]bool{})
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d: %+v", len(routes), routes)
	}
	if routes[0].path != "/api/v1/items" {
		t.Errorf("path = %q, want /api/v1/items", routes[0].path)
	}
}
