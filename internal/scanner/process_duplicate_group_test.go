//ff:func feature=scan type=test control=sequence topic=scanner
//ff:what processDuplicateGroup 테스트 (round5)
package scanner

import "testing"

func TestProcessDuplicateGroup_Round5(t *testing.T) {
	endpoints := []Endpoint{
		{Method: "GET", Path: "/users"},
		{Method: "POST", Path: "/users"},
	}
	result := map[int]string{}
	seen := map[string]bool{}
	processDuplicateGroup("user", []int{0, 1}, endpoints, result, seen)

	if len(result) != 2 {
		t.Fatalf("expected 2 results, got %d: %v", len(result), result)
	}
	// both indices must be assigned distinct, deterministic operation IDs
	if result[0] == result[1] {
		t.Fatalf("expected distinct ids, got %q and %q", result[0], result[1])
	}
	if !seen[result[0]] || !seen[result[1]] {
		t.Fatalf("seen not updated: %v", seen)
	}
}
