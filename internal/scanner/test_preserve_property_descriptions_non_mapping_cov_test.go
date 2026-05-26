//ff:func feature=scan type=test control=sequence
//ff:what TestPreservePropertyDescriptions_NonMappingCov 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestPreservePropertyDescriptions_NonMappingCov(t *testing.T) {
	scanSchema := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "properties"},
		{Kind: yaml.ScalarNode, Value: "notamapping"},
	}}
	baseSchema := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "properties"},
		{Kind: yaml.ScalarNode, Value: "notamapping"},
	}}
	preservePropertyDescriptions(scanSchema, baseSchema)
}
