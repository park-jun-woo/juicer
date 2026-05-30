//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestHasFieldCall_NotField 테스트
package fastapi

import "testing"

func TestHasFieldCall_NotField(t *testing.T) {
	assign, src := firstAssignment(t, []byte("x: int = other()\n"))
	if hasFieldCall(assign, src) {
		t.Fatal("expected false")
	}
}
