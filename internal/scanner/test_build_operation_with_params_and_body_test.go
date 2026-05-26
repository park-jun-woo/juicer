//ff:func feature=scan type=test control=sequence
//ff:what TestBuildOperation_WithParamsAndBody 테스트
package scanner

import "testing"

func TestBuildOperation_WithParamsAndBody(t *testing.T) {
	ep := Endpoint{
		Method: "POST",
		Path:   "/api/users",
		Request: &Request{
			Query: []Param{{Name: "id", Type: "string"}},
			Body:  &Body{Fields: []Field{{Name: "name", JSON: "name", Type: "string"}}},
		},
	}
	schemas := map[string]any{}
	op := buildOperation(ep, schemas)
	if op["parameters"] == nil {
		t.Fatal("expected parameters")
	}
	if op["requestBody"] == nil {
		t.Fatal("expected requestBody")
	}
}
