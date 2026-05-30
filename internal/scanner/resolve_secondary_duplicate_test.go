//ff:func feature=scan type=test control=sequence topic=scanner
//ff:what resolveSecondaryDuplicate 테스트 (round5)
package scanner

import "testing"

func TestResolveSecondaryDuplicate_Round5(t *testing.T) {
	// not seen => returned unchanged
	seen := map[string]bool{}
	if got := resolveSecondaryDuplicate("getUser", seen); got != "getUser" {
		t.Fatalf("unseen: got %q", got)
	}

	// seen once => suffix 2
	seen = map[string]bool{"getUser": true}
	if got := resolveSecondaryDuplicate("getUser", seen); got != "getUser2" {
		t.Fatalf("seen: got %q", got)
	}

	// seen and 2 also taken => 3
	seen = map[string]bool{"getUser": true, "getUser2": true}
	if got := resolveSecondaryDuplicate("getUser", seen); got != "getUser3" {
		t.Fatalf("seen+2: got %q", got)
	}
}
