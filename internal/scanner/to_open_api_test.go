package scanner

import "testing"

func TestToOpenAPI_EmptyResult(t *testing.T) {
	result := &ScanResult{}
	data, err := ToOpenAPI(result)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("expected non-empty")
	}
}

func TestToOpenAPI_WithEndpoints(t *testing.T) {
	result := &ScanResult{
		Endpoints: []Endpoint{
			{Method: "GET", Path: "/api/users", Handler: "h.ListUsers"},
		},
	}
	data, err := ToOpenAPI(result)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("expected non-empty")
	}
}
