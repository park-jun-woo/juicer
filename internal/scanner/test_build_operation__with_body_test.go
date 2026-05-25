//ff:func feature=scan type=convert control=sequence
//ff:what TestBuildOperation_WithBody 테스트
package scanner

import "testing"

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
