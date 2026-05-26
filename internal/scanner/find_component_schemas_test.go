//ff:func feature=scan type=test control=sequence
//ff:what TestFindComponentSchemas_Found 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestFindComponentSchemas_Found(t *testing.T) {
	schemas := &yaml.Node{Kind: yaml.MappingNode}
	comp := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "schemas"},
		schemas,
	}}
	root := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "components"},
		comp,
	}}
	result := findComponentSchemas(root)
	if result != schemas {
		t.Fatal("expected schemas node")
	}
}

