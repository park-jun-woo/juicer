//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestHasFieldCall_True 테스트
package fastapi

import "testing"

func TestHasFieldCall_True(t *testing.T) {
	assign, src := firstAssignment(t, []byte("x: int = Field(ge=0)\n"))
	if !hasFieldCall(assign, src) {
		t.Fatal("expected true")
	}
}
