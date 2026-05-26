//ff:func feature=scan type=test control=sequence
//ff:what TestResponseSchema_SliceCov 테스트
package scanner

import "testing"

func TestResponseSchema_SliceCov(t *testing.T) {
	resp := Response{TypeName: "[]User", Fields: []Field{{Name: "ID", Type: "int"}}}
	schemas := map[string]any{}
	result := responseSchema(resp, schemas)
	if result["type"] != "array" {
		t.Fatal("expected array")
	}
}
