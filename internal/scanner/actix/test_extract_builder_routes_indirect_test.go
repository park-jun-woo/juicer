//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractBuilderRoutes_Indirect 테스트
package actix

import "testing"

// Minimal reproduction of BUG-008: a `fn s() -> Scope` builder chain registered
// indirectly via App.service(s()) must yield the scope-prefixed route.
func TestExtractBuilderRoutes_Indirect(t *testing.T) {
	fi := aFi(t, `fn s() -> Scope { web::scope("/p").service(web::resource("/x").route(web::get().to(h))) }
fn main() { App::new().service(s()); }`)
	handlerFuncs := map[string]*handlerInfo{}
	collectHandlerFuncs(fi, handlerFuncs)

	routes := extractBuilderRoutes(fi, handlerFuncs)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d: %+v", len(routes), routes)
	}
	if routes[0].path != "/p/x" || routes[0].method != "GET" || routes[0].handler != "h" {
		t.Errorf("unexpected route: %+v", routes[0])
	}
}
