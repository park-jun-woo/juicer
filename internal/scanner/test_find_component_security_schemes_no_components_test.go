//ff:func feature=scan type=test control=sequence
//ff:what TestFindComponentSecuritySchemes_NoComponents 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestFindComponentSecuritySchemes_NoComponents(t *testing.T) {
	empty := &yaml.Node{Kind: yaml.MappingNode}
	if findComponentSecuritySchemes(empty) != nil {
		t.Fatal("expected nil for missing components")
	}
}
