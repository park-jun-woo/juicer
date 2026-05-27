//ff:func feature=scan type=test control=sequence
//ff:what TestResponseSchema_EmptyFieldsCov 테스트
package scanner

import "testing"

func TestResponseSchema_EmptyFieldsCov(t *testing.T) {
	resp := Response{TypeName: "SomeType"}
	schemas := map[string]any{}
	result := responseSchema(resp, schemas)
	if result["$ref"] != "#/components/schemas/someType" {
		t.Fatalf("expected $ref to someType, got %v", result)
	}
}
