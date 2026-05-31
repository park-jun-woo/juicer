//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestResolveIndirectServiceArg_Cycle 테스트
package actix

import "testing"

func TestResolveIndirectServiceArg_Cycle(t *testing.T) {
	// a() returns a scope that registers b(), and b() registers a() — a cycle.
	// The visited guard must keep recursion finite.
	fi := aFi(t, `pub fn a() -> Scope { web::scope("/a").service(b()) }
pub fn b() -> Scope { web::scope("/b").service(a()) }
fn cfg(c: &mut ServiceConfig) { c.service(a()); }`)
	handlerFuncs := map[string]*handlerInfo{}
	collectHandlerFuncs(fi, handlerFuncs)

	cfg := findFuncByName(fi.root, fi.src, "cfg")
	arg := findCallByFuncName(cfg, fi.src, "a") // a()

	var routes []builderRoute
	resolveIndirectServiceArg(arg, fi.src, "", &routes, handlerFuncs, map[string]bool{})
	// No concrete web::resource/route inside, so no routes; the point is it
	// terminates without stack overflow.
	if len(routes) != 0 {
		t.Errorf("expected 0 routes, got %+v", routes)
	}
}
