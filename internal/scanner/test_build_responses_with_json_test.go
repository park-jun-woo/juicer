//ff:func feature=scan type=test control=sequence
//ff:what TestBuildResponses_WithJSON 테스트
package scanner

import "testing"

func TestBuildResponses_WithJSON(t *testing.T) {
	resps := []Response{{Status: "200", Kind: "json", Fields: []Field{{Name: "id", JSON: "id", Type: "int"}}}}
	schemas := map[string]any{}
	result := buildResponses(resps, schemas)
	if result["200"] == nil {
		t.Fatal("expected 200 response")
	}
}
