package scanner

import "testing"

func TestBuildOperationParams_Nil(t *testing.T) {
	result := buildOperationParams(nil)
	if result != nil {
		t.Fatal("expected nil")
	}
}

func TestBuildOperationParams_WithParams(t *testing.T) {
	req := &Request{
		PathParams: []Param{{Name: "id"}},
		Query:      []Param{{Name: "page"}},
	}
	result := buildOperationParams(req)
	if len(result) != 2 {
		t.Fatalf("expected 2, got %d", len(result))
	}
}
