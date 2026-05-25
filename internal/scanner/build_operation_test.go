package scanner

import "testing"

func TestBuildOperation_Basic(t *testing.T) {
	ep := Endpoint{Method: "GET", Path: "/api/health"}
	schemas := map[string]any{}
	op := buildOperation(ep, schemas)
	if op == nil {
		t.Fatal("expected non-nil")
	}
}

func TestBuildOperation_WithBody(t *testing.T) {
	ep := Endpoint{
		Method: "POST",
		Path:   "/api/users",
		Request: &Request{
			Body: &Body{TypeName: "User", Fields: []Field{{Name: "name", JSON: "name"}}},
		},
		Responses: []Response{{Status: "201"}},
	}
	schemas := map[string]any{}
	op := buildOperation(ep, schemas)
	if op["requestBody"] == nil {
		t.Fatal("expected requestBody")
	}
}
