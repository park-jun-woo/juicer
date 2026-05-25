//ff:func feature=scan type=extract control=sequence
//ff:what TestBodySchema_Slice 테스트
package scanner

import "testing"

func TestBodySchema_Slice(t *testing.T) {
	body := &Body{TypeName: "[]User", Fields: []Field{{Name: "id", JSON: "id"}}}
	schemas := map[string]any{}
	result := bodySchema(body, schemas)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}
