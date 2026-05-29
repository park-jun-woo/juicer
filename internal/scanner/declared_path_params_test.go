//ff:func feature=scan type=test control=sequence
//ff:what TestDeclaredPathParams 테스트
package scanner

import "testing"

func TestDeclaredPathParams(t *testing.T) {
	params := []map[string]any{
		{"name": "id", "in": "path"},
		{"name": "q", "in": "query"},
		{"name": "slug", "in": "path"},
		{"in": "path"}, // no name -> ignored
	}
	declared := declaredPathParams(params)
	if !declared["id"] || !declared["slug"] {
		t.Fatalf("expected id and slug declared, got %v", declared)
	}
	if declared["q"] {
		t.Fatal("query param should not be in declared path set")
	}
	if len(declared) != 2 {
		t.Fatalf("expected 2 declared path params, got %d", len(declared))
	}
}
