//ff:func feature=scan type=test control=sequence
//ff:what TestBuildResponses_Empty 테스트
package scanner

import "testing"

func TestBuildResponses_Empty(t *testing.T) {
	schemas := map[string]any{}
	result := buildResponses(nil, schemas)
	if result == nil {
		t.Fatal("expected non-nil")
	}

	// with responses
	resps := []Response{
		{Status: "200", Kind: "json", Fields: []Field{{Name: "id", JSON: "id", Type: "int"}}},
		{Status: "200", Kind: "json", TypeName: "User", Fields: []Field{{Name: "name", JSON: "name", Type: "string"}}, Confidence: "partial"},
		{Status: "404", Kind: "text"},
	}
	result = buildResponses(resps, schemas)
	if result["200"] == nil {
		t.Fatal("expected 200 response")
	}
	if result["404"] == nil {
		t.Fatal("expected 404 response")
	}
}
