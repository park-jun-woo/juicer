//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestCollectParentNames_WithParents 테스트
package fastapi

import "testing"

func TestCollectParentNames_WithParents(t *testing.T) {
	cls, src := firstClass(t, []byte("class User(Base, BaseModel, table=True): pass\n"))
	names := collectParentNames(cls, src)
	if len(names) != 2 || names[0] != "Base" || names[1] != "BaseModel" {
		t.Fatalf("got %v", names)
	}
}
