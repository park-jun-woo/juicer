//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestHasFieldCall_NoCall 테스트
package fastapi

import "testing"

func TestHasFieldCall_NoCall(t *testing.T) {
	assign, src := firstAssignment(t, []byte("x: int = 5\n"))
	if hasFieldCall(assign, src) {
		t.Fatal("expected false")
	}
}
