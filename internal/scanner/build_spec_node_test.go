package scanner

import "testing"

func TestBuildSpecNode_Empty(t *testing.T) {
	result := &ScanResult{Endpoints: nil}
	node := buildSpecNode(result)
	if node == nil {
		t.Fatal("expected non-nil")
	}
}

func TestBuildSpecNode_WithEndpoints(t *testing.T) {
	result := &ScanResult{
		Endpoints: []Endpoint{
			{Method: "GET", Path: "/api/health", Responses: []Response{{Status: "200"}}},
		},
	}
	node := buildSpecNode(result)
	if node == nil {
		t.Fatal("expected non-nil")
	}
}
