//ff:func feature=scan type=test control=sequence
//ff:what TestBuildOperation_Basic 테스트
package scanner

import "testing"

func TestBuildOperation_Basic(t *testing.T) {
	ep := Endpoint{Method: "GET", Path: "/api/health"}
	schemas := map[string]any{}
	op := buildOperation(ep, schemas)
	if op == nil {
		t.Fatal("expected non-nil")
	}

	// with params and request body
	ep2 := Endpoint{
		Method: "POST",
		Path:   "/api/users/:id",
		Request: &Request{
			PathParams: []Param{{Name: "id", Type: "string"}},
			Query:      []Param{{Name: "q", Type: "string"}},
			Body:       &Body{Fields: []Field{{Name: "name", JSON: "name", Type: "string"}}},
		},
		Responses: []Response{{Status: "200"}},
	}
	op2 := buildOperation(ep2, schemas)
	if op2["parameters"] == nil {
		t.Fatal("expected parameters")
	}
	if op2["requestBody"] == nil {
		t.Fatal("expected requestBody")
	}
}
