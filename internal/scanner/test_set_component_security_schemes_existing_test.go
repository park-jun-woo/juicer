//ff:func feature=scan type=test control=sequence
//ff:what TestSetComponentSecuritySchemes_ExistingComponents 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestSetComponentSecuritySchemes_ExistingComponents(t *testing.T) {
	comp := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "schemas"},
		{Kind: yaml.MappingNode},
	}}
	node := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "components"},
		comp,
	}}
	schemes := &yaml.Node{Kind: yaml.MappingNode}
	setComponentSecuritySchemes(node, schemes)
	result := findMappingValue(comp, "securitySchemes")
	if result != schemes {
		t.Fatal("expected securitySchemes added to existing components")
	}
	// schemas should still be there
	if findMappingValue(comp, "schemas") == nil {
		t.Fatal("expected schemas preserved")
	}
}
