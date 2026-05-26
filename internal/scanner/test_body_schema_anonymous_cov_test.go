//ff:func feature=scan type=test control=sequence
//ff:what TestBodySchema_AnonymousCov 테스트
package scanner

import "testing"

func TestBodySchema_AnonymousCov(t *testing.T) {
	body := &Body{Fields: []Field{{Name: "id", JSON: "id", Type: "int"}}}
	schemas := map[string]any{}
	result := bodySchema(body, schemas)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}
