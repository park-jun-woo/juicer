//ff:func feature=scan type=test control=sequence
//ff:what TestPreservePropertyDescriptions_NoMatchCov 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestPreservePropertyDescriptions_NoMatchCov(t *testing.T) {
	scanSchema := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "properties"},
		{Kind: yaml.MappingNode, Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "email"},
			{Kind: yaml.MappingNode},
		}},
	}}
	baseSchema := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "properties"},
		{Kind: yaml.MappingNode},
	}}
	preservePropertyDescriptions(scanSchema, baseSchema)
}
