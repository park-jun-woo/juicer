//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestIntPtr 테스트
package spring

import "testing"

func TestIntPtr(t *testing.T) {
	if p := intPtr(9); p == nil || *p != 9 {
		t.Fatalf("got %v", p)
	}
}
