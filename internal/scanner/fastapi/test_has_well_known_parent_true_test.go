//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestHasWellKnownParent_True 테스트
package fastapi

import "testing"

func TestHasWellKnownParent_True(t *testing.T) {
	cls, src := firstClass(t, []byte("class User(BaseModel):\n    id: int\n"))
	if !hasWellKnownParent(cls, src) {
		t.Fatal("expected true")
	}
}
