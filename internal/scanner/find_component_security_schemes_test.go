//ff:func feature=scan type=test control=sequence
//ff:what TestFindComponentSecuritySchemes_Found 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestFindComponentSecuritySchemes_Found(t *testing.T) {
	schemes := &yaml.Node{Kind: yaml.MappingNode}
	comp := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "securitySchemes"},
		schemes,
	}}
	root := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "components"},
		comp,
	}}
	result := findComponentSecuritySchemes(root)
	if result != schemes {
		t.Fatal("expected securitySchemes node")
	}
}
