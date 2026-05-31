//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFuncReturnScopeChain_None 테스트
package actix

import "testing"

func TestFuncReturnScopeChain_None(t *testing.T) {
	root, src := aParse(t, `fn plain() -> i32 { 1 + 2 }`)
	fn := findFuncByName(root, src, "plain")
	if fn == nil {
		t.Fatal("func plain not found")
	}
	if chain := funcReturnScopeChain(fn, src); chain != nil {
		t.Errorf("expected nil, got %q", nodeText(chain, src))
	}
}
