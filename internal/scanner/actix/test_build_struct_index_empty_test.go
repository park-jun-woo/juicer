//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestBuildStructIndex_Empty 테스트
package actix

import "testing"

func TestBuildStructIndex_Empty(t *testing.T) {
	idx := buildStructIndex(nil)
	if len(idx) != 0 {
		t.Fatalf("expected empty index, got %d", len(idx))
	}
}
