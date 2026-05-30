//ff:func feature=scan type=test control=sequence topic=express
//ff:what hasNeededName: 미해결 포함 true / 미포함 false
package express

import "testing"

func TestHasNeededName(t *testing.T) {
	set := map[string]bool{"A": true}
	if !hasNeededName([]string{"X", "A"}, set) {
		t.Error("expected true when A unresolved")
	}
	if hasNeededName([]string{"X", "Y"}, set) {
		t.Error("expected false when none unresolved")
	}
	if hasNeededName(nil, set) {
		t.Error("expected false for empty names")
	}
}
