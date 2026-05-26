//ff:func feature=scan type=test control=sequence
//ff:what TestResponseSchema_EmptyFieldsCov 테스트
package scanner

import "testing"

func TestResponseSchema_EmptyFieldsCov(t *testing.T) {
	resp := Response{TypeName: "SomeType"}
	schemas := map[string]any{}
	result := responseSchema(resp, schemas)
	if result["type"] != "object" {
		t.Fatal("expected object")
	}
}
