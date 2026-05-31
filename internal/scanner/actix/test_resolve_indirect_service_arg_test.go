//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestResolveIndirectServiceArg 테스트
package actix

import "testing"

func TestResolveIndirectServiceArg(t *testing.T) {
	fi := aFi(t, `pub fn s() -> Scope { web::scope("/p").service(web::resource("/x").route(web::get().to(h))) }
fn cfg(c: &mut ServiceConfig) { c.service(s()); }`)
	handlerFuncs := map[string]*handlerInfo{}
	collectHandlerFuncs(fi, handlerFuncs)

	// Locate the cfg function body's `s()` call (the indirect registration).
	cfg := findFuncByName(fi.root, fi.src, "cfg")
	arg := findCallByFuncName(cfg, fi.src, "s") // s()

	var routes []builderRoute
	resolveIndirectServiceArg(arg, fi.src, "/api", &routes, handlerFuncs, map[string]bool{})
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d: %+v", len(routes), routes)
	}
	if routes[0].path != "/api/p/x" {
		t.Errorf("path = %q, want /api/p/x", routes[0].path)
	}
	if routes[0].method != "GET" || routes[0].handler != "h" {
		t.Errorf("unexpected route: %+v", routes[0])
	}
}
