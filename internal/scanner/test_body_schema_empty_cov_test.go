//ff:func feature=scan type=test control=sequence
//ff:what TestBodySchema_EmptyCov 테스트
package scanner

import "testing"

func TestBodySchema_EmptyCov(t *testing.T) {
	body := &Body{}
	schemas := map[string]any{}
	result := bodySchema(body, schemas)
	if result["type"] != "object" {
		t.Fatalf("expected object type, got %v", result)
	}
}
