//ff:func feature=scan type=test control=sequence
//ff:what TestBuildOperation_WithRequestCov 테스트
package scanner

import "testing"

func TestBuildOperation_WithRequestCov(t *testing.T) {
	ep := Endpoint{
		Method: "POST",
		Path:   "/api/users/:id",
		Request: &Request{
			PathParams: []Param{{Name: "id", Type: "int"}},
			Body:       &Body{Fields: []Field{{Name: "name", JSON: "name", Type: "string"}}},
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
