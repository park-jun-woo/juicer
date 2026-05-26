//ff:func feature=scan type=test control=sequence
//ff:what TestResponseSchema_NoTypeCov 테스트
package scanner

import "testing"

func TestResponseSchema_NoTypeCov(t *testing.T) {
	resp := Response{}
	schemas := map[string]any{}
	result := responseSchema(resp, schemas)
	if result["type"] != "object" {
		t.Fatal("expected object")
	}
}
