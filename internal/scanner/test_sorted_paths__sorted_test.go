//ff:func feature=scan type=extract control=sequence
//ff:what TestSortedPaths_Sorted 테스트
package scanner

import "testing"

func TestSortedPaths_Sorted(t *testing.T) {
	paths := map[string]map[string]any{
		"/b": {"get": map[string]any{}},
		"/a": {"post": map[string]any{}},
	}
	node := sortedPaths(paths)
	if node.Content[0].Value != "/a" {
		t.Fatal("expected /a first")
	}
}
