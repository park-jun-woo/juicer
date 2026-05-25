package scanner

import "testing"

func TestDeduplicateEndpoints_Empty(t *testing.T) {
	result := deduplicateEndpoints(nil)
	if len(result) != 0 {
		t.Fatal("expected empty")
	}
}

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
