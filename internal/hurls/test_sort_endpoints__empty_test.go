//ff:func feature=hurl type=parse control=sequence
//ff:what TestSortEndpoints_Empty 테스트
package hurls

import "testing"

func TestSortEndpoints_Empty(t *testing.T) {
	sorted := sortEndpoints(nil)
	if len(sorted) != 0 {
		t.Fatalf("expected 0, got %d", len(sorted))
	}
}
