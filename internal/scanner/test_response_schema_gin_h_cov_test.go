//ff:func feature=scan type=test control=sequence
//ff:what TestResponseSchema_GinHCov 테스트
package scanner

import "testing"

func TestResponseSchema_GinHCov(t *testing.T) {
	resp := Response{TypeName: "gin.H", Fields: []Field{{Name: "msg", JSON: "msg", Type: "string"}}}
	schemas := map[string]any{}
	result := responseSchema(resp, schemas)
	if result["type"] != "object" {
		t.Fatalf("expected object, got %v", result["type"])
	}
}
