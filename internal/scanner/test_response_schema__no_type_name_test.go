//ff:func feature=scan type=extract control=sequence
//ff:what TestResponseSchema_NoTypeName 테스트
package scanner

import "testing"

func TestResponseSchema_NoTypeName(t *testing.T) {
	resp := Response{}
	schemas := map[string]any{}
	result := responseSchema(resp, schemas)
	if result["type"] != "object" {
		t.Fatal("expected object")
	}
}
