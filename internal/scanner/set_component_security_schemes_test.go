//ff:func feature=scan type=test control=sequence
//ff:what TestSetComponentSecuritySchemes_New 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestSetComponentSecuritySchemes_New(t *testing.T) {
	node := &yaml.Node{Kind: yaml.MappingNode}
	schemes := &yaml.Node{Kind: yaml.MappingNode}
	setComponentSecuritySchemes(node, schemes)
	comp := findMappingValue(node, "components")
	if comp == nil {
		t.Fatal("expected components")
	}
	result := findMappingValue(comp, "securitySchemes")
	if result != schemes {
		t.Fatal("expected securitySchemes")
	}
}
