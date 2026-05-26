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

	// path param uses its own type, not hardcoded "string"
	schema0 := params[0]["schema"].(map[string]any)
	if schema0["type"] != "string" {
		t.Fatalf("path param type: got %v, want string", schema0["type"])
	}

	// typed path param preserves type
	reqTyped := &Request{
		PathParams: []Param{{Name: "id", Type: "string:uuid"}},
	}
	paramsTyped := buildOperationParams(reqTyped)
	schemaTyped := paramsTyped[0]["schema"].(map[string]any)
	if schemaTyped["type"] != "string" {
		t.Fatalf("typed path type: got %v, want string", schemaTyped["type"])
	}
	if schemaTyped["format"] != "uuid" {
		t.Fatalf("typed path format: got %v, want uuid", schemaTyped["format"])
	}
}
