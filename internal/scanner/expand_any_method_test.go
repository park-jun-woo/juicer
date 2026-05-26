//ff:func feature=scan type=test control=sequence
//ff:what TestExpandAnyMethod 테스트
package scanner

import "testing"

func TestExpandAnyMethod(t *testing.T) {
	got := expandAnyMethod("any")
	if len(got) != 5 {
		t.Fatalf("expected 5 methods for 'any', got %d", len(got))
	}
	got2 := expandAnyMethod("get")
	if len(got2) != 1 || got2[0] != "get" {
		t.Fatalf("expected [get] for 'get', got %v", got2)
	}
}
