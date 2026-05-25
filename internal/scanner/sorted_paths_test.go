//ff:func feature=scan type=extract control=sequence
//ff:what TestSortedPaths_Empty 테스트
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
