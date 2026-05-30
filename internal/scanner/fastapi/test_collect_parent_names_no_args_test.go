//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestCollectParentNames_NoArgs 테스트
package fastapi

import "testing"

func TestCollectParentNames_NoArgs(t *testing.T) {
	cls, src := firstClass(t, []byte("class Plain: pass\n"))
	if names := collectParentNames(cls, src); names != nil {
		t.Fatalf("expected nil, got %v", names)
	}
}
