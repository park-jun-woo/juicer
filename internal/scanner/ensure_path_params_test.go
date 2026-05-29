//ff:func feature=scan type=test control=sequence
//ff:what TestEnsurePathParams 테스트
package scanner

import "testing"

func TestEnsurePathParams(t *testing.T) {
	// already-declared path param is preserved (not duplicated)
	existing := []map[string]any{
		{"name": "id", "in": "path", "required": true, "schema": buildParamSchema("integer")},
	}
	got := ensurePathParams(existing, "/users/{id}/posts/{slug}")
	if len(got) != 2 {
		t.Fatalf("expected 2 params, got %d", len(got))
	}
	// existing typed param kept
	if got[0]["schema"].(map[string]any)["type"] != "integer" {
		t.Fatalf("existing typed path param not preserved")
	}
	// missing token auto-declared as string in:path
	if got[1]["name"] != "slug" || got[1]["in"] != "path" {
		t.Fatalf("missing token not declared: %v", got[1])
	}
	if got[1]["schema"].(map[string]any)["type"] != "string" {
		t.Fatalf("auto-declared path param should be string")
	}

	// no tokens -> unchanged nil
	if ensurePathParams(nil, "/health") != nil {
		t.Fatal("expected nil for tokenless path")
	}
}
