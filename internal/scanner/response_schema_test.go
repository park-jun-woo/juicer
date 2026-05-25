//ff:func feature=scan type=extract control=sequence
//ff:what TestResponseSchema_NamedType 테스트
package scanner

import "testing"

func TestResponseSchema_NamedType(t *testing.T) {
	resp := Response{TypeName: "User", Fields: []Field{{Name: "ID", Type: "int"}}}
	schemas := map[string]any{}
	result := responseSchema(resp, schemas)
	if result["$ref"] == nil {
		t.Fatal("expected $ref")
	}
}
