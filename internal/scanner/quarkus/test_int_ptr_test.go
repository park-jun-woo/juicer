//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestIntPtr 테스트
package quarkus

import "testing"

func TestIntPtr(t *testing.T) {
	p := intPtr(7)
	if p == nil || *p != 7 {
		t.Fatalf("got %v", p)
	}
}
