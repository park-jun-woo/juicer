//ff:func feature=scan type=test control=sequence
//ff:what TestMergePathItem 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergePathItem(t *testing.T) {
	scanOps := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "get"},
		{Kind: yaml.MappingNode},
	}}
	baseOps := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "get"},
		{Kind: yaml.MappingNode},
	}}
	result := mergePathItem(scanOps, baseOps)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}
