//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what expandAllMethod 테스트
package fastify

import "testing"

func TestExpandAllMethod(t *testing.T) {
	all := expandAllMethod("all")
	if len(all) != 5 {
		t.Fatalf("expected 5 methods for 'all', got %d", len(all))
	}
	single := expandAllMethod("GET")
	if len(single) != 1 || single[0] != "GET" {
		t.Fatalf("expected [GET], got %v", single)
	}
}
