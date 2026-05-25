//ff:func feature=scan type=extract control=sequence
//ff:what TestDeduplicateEndpoints_NoDuplicates 테스트
package scanner

import "testing"

func TestDeduplicateEndpoints_NoDuplicates(t *testing.T) {
	eps := []Endpoint{
		{Method: "GET", Path: "/a"},
		{Method: "POST", Path: "/b"},
	}
	result := deduplicateEndpoints(eps)
	if len(result) != 2 {
		t.Fatalf("expected 2, got %d", len(result))
	}
}
