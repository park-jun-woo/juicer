//ff:func feature=scan type=test control=sequence
//ff:what TestFindMappingIndex_Found 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestFindMappingIndex_Found(t *testing.T) {
	node := &yaml.Node{
		Kind: yaml.MappingNode,
		Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "key1"},
			{Kind: yaml.ScalarNode, Value: "val1"},
			{Kind: yaml.ScalarNode, Value: "key2"},
			{Kind: yaml.ScalarNode, Value: "val2"},
		},
	}
	if idx := findMappingIndex(node, "key2"); idx != 2 {
		t.Fatalf("expected 2, got %d", idx)
	}
}
