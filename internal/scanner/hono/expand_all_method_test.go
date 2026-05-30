//ff:func feature=scan type=test control=sequence topic=hono
//ff:what expandAllMethod 테스트
package hono

import "testing"

func TestExpandAllMethod(t *testing.T) {
	if got := expandAllMethod("all"); len(got) != 5 {
		t.Fatalf("expected 5 for 'all', got %d", len(got))
	}
	if got := expandAllMethod("GET"); len(got) != 1 || got[0] != "GET" {
		t.Fatalf("expected [GET], got %v", got)
	}
}
