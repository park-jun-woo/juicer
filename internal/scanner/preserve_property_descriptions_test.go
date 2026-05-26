//ff:func feature=scan type=test control=sequence
//ff:what TestPreservePropertyDescriptions 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestPreservePropertyDescriptions(t *testing.T) {
	scanSchema := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "properties"},
		{Kind: yaml.MappingNode, Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "name"},
			{Kind: yaml.MappingNode},
		}},
	}}
	baseSchema := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "properties"},
		{Kind: yaml.MappingNode, Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "name"},
			{Kind: yaml.MappingNode, Content: []*yaml.Node{
				{Kind: yaml.ScalarNode, Value: "description"},
				{Kind: yaml.ScalarNode, Value: "User name"},
			}},
		}},
	}}
	preservePropertyDescriptions(scanSchema, baseSchema)
}
