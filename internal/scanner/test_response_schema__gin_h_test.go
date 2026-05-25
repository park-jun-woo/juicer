//ff:func feature=scan type=extract control=sequence
//ff:what TestResponseSchema_GinH 테스트
package scanner

import "testing"

func TestResponseSchema_GinH(t *testing.T) {
	resp := Response{TypeName: "gin.H", Fields: []Field{{Name: "msg", Type: "string"}}}
	schemas := map[string]any{}
	result := responseSchema(resp, schemas)
	if result["type"] != "object" {
		t.Fatalf("expected object, got %v", result["type"])
	}
}
