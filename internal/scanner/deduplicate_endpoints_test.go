//ff:func feature=scan type=test control=sequence
//ff:what TestDeduplicateEndpoints_Empty 테스트
package scanner

import "testing"

func TestDeduplicateEndpoints_Empty(t *testing.T) {
	result := deduplicateEndpoints(nil)
	if len(result) != 0 {
		t.Fatal("expected empty")
	}

	// with duplicates (richer one wins)
	eps := []Endpoint{
		{Method: "GET", Path: "/api/users"},
		{Method: "GET", Path: "/api/users", Responses: []Response{{Status: "200", Kind: "json", Fields: []Field{{Name: "id"}}}}},
		{Method: "POST", Path: "/api/users"},
	}
	result = deduplicateEndpoints(eps)
	if len(result) != 2 {
		t.Fatalf("expected 2 deduplicated, got %d", len(result))
	}
	// The richer one should win (the one with responses)
	if len(result[0].Responses) == 0 {
		t.Fatal("expected the richer endpoint to win")
	}
}

