//ff:func feature=scan type=convert control=sequence
//ff:what TestSortedPaths 테스트
package scanner

import (
	"testing"
)

func TestSortedPaths(t *testing.T) {
	paths := map[string]map[string]any{
		"/z": {"get": map[string]any{}},
		"/a": {"post": map[string]any{}},
	}
	node := sortedPaths(paths)
	if node == nil {
		t.Fatal("expected non-nil node")
	}
	// First key should be /a
	if node.Content[0].Value != "/a" {
		t.Errorf("expected first path /a, got %q", node.Content[0].Value)
	}
}
