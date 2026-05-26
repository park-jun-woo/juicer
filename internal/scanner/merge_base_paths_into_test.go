//ff:func feature=scan type=test control=sequence
//ff:what TestMergeBasePathsInto 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeBasePathsInto(t *testing.T) {
	merged := &yaml.Node{Kind: yaml.MappingNode}
	basePaths := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "/api/users"},
		{Kind: yaml.MappingNode, Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "get"},
			{Kind: yaml.MappingNode},
		}},
	}}
	reg := map[string]bool{"get\t/api/users": true}
	mergeBasePathsInto(merged, basePaths, reg)
	if len(merged.Content) != 2 {
		t.Fatalf("expected 2 nodes, got %d", len(merged.Content))
	}
}
