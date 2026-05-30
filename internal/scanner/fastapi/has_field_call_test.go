//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what hasFieldCall: Field() true / 비Field / call없음 분기
package fastapi

import "testing"

func TestHasFieldCall_True(t *testing.T) {
	assign, src := firstAssignment(t, []byte("x: int = Field(ge=0)\n"))
	if !hasFieldCall(assign, src) {
		t.Fatal("expected true")
	}
}

func TestHasFieldCall_NotField(t *testing.T) {
	assign, src := firstAssignment(t, []byte("x: int = other()\n"))
	if hasFieldCall(assign, src) {
		t.Fatal("expected false")
	}
}

func TestHasFieldCall_NoCall(t *testing.T) {
	assign, src := firstAssignment(t, []byte("x: int = 5\n"))
	if hasFieldCall(assign, src) {
		t.Fatal("expected false")
	}
}
