//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what hasWellKnownParent: BaseModel 부모 true / 일반 부모 false
package fastapi

import "testing"

func TestHasWellKnownParent_True(t *testing.T) {
	cls, src := firstClass(t, []byte("class User(BaseModel):\n    id: int\n"))
	if !hasWellKnownParent(cls, src) {
		t.Fatal("expected true")
	}
}

func TestHasWellKnownParent_False(t *testing.T) {
	cls, src := firstClass(t, []byte("class User(SomethingElse):\n    id: int\n"))
	if hasWellKnownParent(cls, src) {
		t.Fatal("expected false")
	}
}
