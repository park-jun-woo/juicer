//ff:func feature=scan type=test control=sequence
//ff:what TestPreserveSchemaDescription_Both 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestPreserveSchemaDescription_Both(t *testing.T) {
	scan := &yaml.Node{Kind: yaml.MappingNode}
	base := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "description"},
		{Kind: yaml.ScalarNode, Value: "Base desc"},
	}}
	preserveSchemaDescription(scan, base)
}

