//ff:func feature=scan type=extract control=sequence
//ff:what TestBodySchema_NoFields 테스트
package scanner

import "testing"

func TestBodySchema_NoFields(t *testing.T) {
	body := &Body{}
	schemas := map[string]any{}
	result := bodySchema(body, schemas)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}
