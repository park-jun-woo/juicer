//ff:func feature=scan type=test control=sequence
//ff:what TestMergeBaseOpsInto 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeBaseOpsInto(t *testing.T) {
	result := &yaml.Node{Kind: yaml.MappingNode}
	baseOps := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "get"},
		{Kind: yaml.MappingNode},
	}}
	scanOps := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "get"},
		{Kind: yaml.MappingNode},
	}}
	added := mergeBaseOpsInto(result, baseOps, scanOps)
	if !added["get"] {
		t.Fatal("expected get to be added")
	}
}
