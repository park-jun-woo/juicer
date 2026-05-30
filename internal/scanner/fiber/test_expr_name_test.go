//ff:func feature=scan type=test control=sequence
//ff:what TestExprName 테스트
package fiber

import "testing"

func TestExprName(t *testing.T) {
	if got := exprNameFor(t, "handler"); got != "handler" {
		t.Errorf("ident: %q", got)
	}
	if got := exprNameFor(t, "pkg.Handler"); got != "pkg.Handler" {
		t.Errorf("selector with recv: %q", got)
	}
	if got := exprNameFor(t, "func() {}"); got != "(inline)" {
		t.Errorf("funclit: %q", got)
	}
	if got := exprNameFor(t, "make()"); got != "make()" {
		t.Errorf("call: %q", got)
	}

	if got := exprNameFor(t, "1 + 2"); got != "" {
		t.Errorf("default: %q, want empty", got)
	}
}
