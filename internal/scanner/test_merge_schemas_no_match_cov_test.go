//ff:func feature=scan type=test control=sequence
//ff:what TestMergeSchemas_NoMatchCov 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeSchemas_NoMatchCov(t *testing.T) {
	scan := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "Order"},
		{Kind: yaml.MappingNode},
	}}
	base := &yaml.Node{Kind: yaml.MappingNode}
	mergeSchemas(scan, base)
}
