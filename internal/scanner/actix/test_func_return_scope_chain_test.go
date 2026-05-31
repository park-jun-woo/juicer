//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFuncReturnScopeChain 테스트
package actix

import "testing"

func TestFuncReturnScopeChain(t *testing.T) {
	root, src := aParse(t, `pub fn s() -> Scope { web::scope("/p").service(web::resource("/x").route(web::get().to(h))) }`)
	fn := findFuncByName(root, src, "s")
	if fn == nil {
		t.Fatal("func s not found")
	}
	chain := funcReturnScopeChain(fn, src)
	if chain == nil {
		t.Fatal("expected scope chain, got nil")
	}
	if root := findCallRoot(chain, src); root != "web::scope" {
		t.Errorf("chain root = %q, want web::scope", root)
	}
}
