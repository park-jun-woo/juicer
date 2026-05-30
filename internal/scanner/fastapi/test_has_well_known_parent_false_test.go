//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestHasWellKnownParent_False 테스트
package fastapi

import "testing"

func TestHasWellKnownParent_False(t *testing.T) {
	cls, src := firstClass(t, []byte("class User(SomethingElse):\n    id: int\n"))
	if hasWellKnownParent(cls, src) {
		t.Fatal("expected false")
	}
}
