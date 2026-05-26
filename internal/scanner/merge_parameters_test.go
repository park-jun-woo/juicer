//ff:func feature=scan type=test control=sequence
//ff:what TestMergeParameters_BothPresent 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeParameters_BothPresent(t *testing.T) {
	scanParams := &yaml.Node{Kind: yaml.SequenceNode, Content: []*yaml.Node{
		{Kind: yaml.MappingNode, Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "name"},
			{Kind: yaml.ScalarNode, Value: "id"},
		}},
	}}
	baseParams := &yaml.Node{Kind: yaml.SequenceNode, Content: []*yaml.Node{
		{Kind: yaml.MappingNode, Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "name"},
			{Kind: yaml.ScalarNode, Value: "id"},
			{Kind: yaml.ScalarNode, Value: "description"},
			{Kind: yaml.ScalarNode, Value: "User ID"},
		}},
	}}
	result := mergeParameters(scanParams, baseParams)
	if result != scanParams {
		t.Fatal("expected scan params returned")
	}
}
