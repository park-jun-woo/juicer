package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestSortedPaths_Empty(t *testing.T) {
	node := sortedPaths(map[string]map[string]any{})
	if node.Kind != yaml.MappingNode {
		t.Fatal("expected mapping node")
	}
	if len(node.Content) != 0 {
		t.Fatal("expected empty")
	}
}

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
