//ff:func feature=scan type=test control=sequence
//ff:what TestPreserveParamDescription_ReplaceDesc 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestPreserveParamDescription_ReplaceDesc(t *testing.T) {
	scan := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "description"},
		{Kind: yaml.ScalarNode, Value: "Old desc"},
	}}
	base := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "description"},
		{Kind: yaml.ScalarNode, Value: "Base desc"},
	}}
	preserveParamDescription(scan, base)
}
