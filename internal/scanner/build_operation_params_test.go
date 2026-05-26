//ff:func feature=scan type=test control=sequence
//ff:what TestBuildOperationParams_Nil 테스트
package scanner

import "testing"

func TestBuildOperationParams_Nil(t *testing.T) {
	result := buildOperationParams(nil)
	if result != nil {
		t.Fatal("expected nil")
	}

	// with path params and query params
	req := &Request{
		PathParams: []Param{{Name: "id", Type: "string"}},
		Query:      []Param{{Name: "q", Type: "string"}, {Name: "page", Type: "int", Default: "1"}},
	}
	params := buildOperationParams(req)
	if len(params) != 3 {
		t.Fatalf("expected 3 params, got %d", len(params))
	}
	if params[0]["in"] != "path" {
		t.Fatal("expected path param")
	}
	if params[1]["in"] != "query" {
		t.Fatal("expected query param")
	}
}
