//ff:func feature=scan type=test control=sequence
//ff:what TestMergeResponses_NoMatchCov 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeResponses_NoMatchCov(t *testing.T) {
	scan := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "404"},
		{Kind: yaml.MappingNode},
	}}
	base := &yaml.Node{Kind: yaml.MappingNode}
	mergeResponses(scan, base)
}
