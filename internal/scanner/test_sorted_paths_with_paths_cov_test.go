//ff:func feature=scan type=test control=sequence
//ff:what TestSortedPaths_WithPathsCov 테스트
package scanner

import "testing"

func TestSortedPaths_WithPathsCov(t *testing.T) {
	paths := map[string]map[string]any{
		"/b": {"get": map[string]any{}},
		"/a": {"post": map[string]any{}},
	}
	node := sortedPaths(paths)
	if len(node.Content) != 4 {
		t.Fatalf("expected 4 nodes, got %d", len(node.Content))
	}
	if node.Content[0].Value != "/a" {
		t.Fatal("expected /a first")
	}
}
