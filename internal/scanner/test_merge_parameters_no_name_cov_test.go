//ff:func feature=scan type=test control=sequence
//ff:what TestMergeParameters_NoNameCov 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeParameters_NoNameCov(t *testing.T) {
	scanParams := &yaml.Node{Kind: yaml.SequenceNode, Content: []*yaml.Node{
		{Kind: yaml.MappingNode, Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "in"},
			{Kind: yaml.ScalarNode, Value: "query"},
		}},
	}}
	baseParams := &yaml.Node{Kind: yaml.SequenceNode}
	mergeParameters(scanParams, baseParams)
}
