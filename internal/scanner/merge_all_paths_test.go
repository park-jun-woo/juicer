//ff:func feature=scan type=test control=sequence
//ff:what TestMergeAllPaths 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMergeAllPaths(t *testing.T) {
	basePaths := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "/api/users"},
		{Kind: yaml.MappingNode, Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "get"},
			{Kind: yaml.MappingNode},
		}},
	}}
	scanPaths := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "/api/health"},
		{Kind: yaml.MappingNode},
	}}
	reg := map[string]bool{"get\t/api/users": true}
	result := mergeAllPaths(basePaths, scanPaths, reg)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}
