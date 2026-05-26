//ff:func feature=scan type=test control=sequence
//ff:what TestSetComponentSchemas_Existing 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestSetComponentSchemas_Existing(t *testing.T) {
	comp := &yaml.Node{Kind: yaml.MappingNode}
	node := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "components"},
		comp,
	}}
	schemas := &yaml.Node{Kind: yaml.MappingNode}
	setComponentSchemas(node, schemas)
}
