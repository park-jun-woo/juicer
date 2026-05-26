//ff:func feature=scan type=test control=sequence
//ff:what TestBodySchema_SliceCov 테스트
package scanner

import "testing"

func TestBodySchema_SliceCov(t *testing.T) {
	body := &Body{TypeName: "[]User", Fields: []Field{{Name: "id", JSON: "id", Type: "int"}}}
	schemas := map[string]any{}
	result := bodySchema(body, schemas)
	if result["type"] != "array" {
		t.Fatal("expected array type")
	}
}
