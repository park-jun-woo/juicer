//ff:func feature=scan type=extract control=sequence
//ff:what TestDeduplicateEndpoints_WithDuplicate 테스트
package scanner

import "testing"

func TestDeduplicateEndpoints_WithDuplicate(t *testing.T) {
	eps := []Endpoint{
		{Method: "GET", Path: "/a"},
		{Method: "GET", Path: "/a", Responses: []Response{{Status: "200"}}},
	}
	result := deduplicateEndpoints(eps)
	if len(result) != 1 {
		t.Fatalf("expected 1, got %d", len(result))
	}
}
