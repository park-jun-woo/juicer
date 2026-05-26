//ff:func feature=scan type=test control=sequence
//ff:what TestDeduplicateEndpoints_WithDuplicatesCov 테스트
package scanner

import "testing"

func TestDeduplicateEndpoints_WithDuplicatesCov(t *testing.T) {
	eps := []Endpoint{
		{Method: "GET", Path: "/api/users"},
		{Method: "GET", Path: "/api/users", Responses: []Response{{Status: "200", Kind: "json", Fields: []Field{{Name: "id"}}}}},
		{Method: "POST", Path: "/api/users"},
	}
	result := deduplicateEndpoints(eps)
	if len(result) != 2 {
		t.Fatalf("expected 2, got %d", len(result))
	}
	// The richer one should win
	if len(result[0].Responses) == 0 {
		t.Fatal("expected the richer endpoint to be kept")
	}
}
