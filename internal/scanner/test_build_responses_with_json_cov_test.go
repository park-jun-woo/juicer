//ff:func feature=scan type=test control=sequence
//ff:what TestBuildResponses_WithJSONCov 테스트
package scanner

import "testing"

func TestBuildResponses_WithJSONCov(t *testing.T) {
	resps := []Response{
		{Status: "200", Kind: "json", Fields: []Field{{Name: "id", JSON: "id", Type: "int"}}},
		{Status: "404", Kind: "json", Fields: []Field{{Name: "error", JSON: "error", Type: "string"}}, Confidence: "partial"},
	}
	schemas := map[string]any{}
	result := buildResponses(resps, schemas)
	if result["200"] == nil {
		t.Fatal("expected 200 response")
	}
	if result["404"] == nil {
		t.Fatal("expected 404 response")
	}
}
