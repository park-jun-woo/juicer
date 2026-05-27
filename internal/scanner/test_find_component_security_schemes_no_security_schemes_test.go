//ff:func feature=scan type=test control=sequence
//ff:what TestFindComponentSecuritySchemes_NoSecuritySchemes 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestFindComponentSecuritySchemes_NoSecuritySchemes(t *testing.T) {
	comp := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "schemas"},
		{Kind: yaml.MappingNode},
	}}
	root := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "components"},
		comp,
	}}
	if findComponentSecuritySchemes(root) != nil {
		t.Fatal("expected nil when no securitySchemes")
	}
}
