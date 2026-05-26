//ff:func feature=scan type=test control=sequence
//ff:what TestBodySchema_EmptyBranch 테스트
package scanner

import "testing"

func TestBodySchema_EmptyBranch(t *testing.T) {
	body := &Body{}
	schemas := map[string]any{}
	result := bodySchema(body, schemas)
	if result["type"] != "object" {
		t.Fatal("expected object type")
	}
}
