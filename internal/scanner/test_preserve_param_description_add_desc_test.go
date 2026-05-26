//ff:func feature=scan type=test control=sequence
//ff:what TestPreserveParamDescription_AddDesc 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestPreserveParamDescription_AddDesc(t *testing.T) {
	scan := &yaml.Node{Kind: yaml.MappingNode}
	base := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "description"},
		{Kind: yaml.ScalarNode, Value: "Param desc"},
	}}
	preserveParamDescription(scan, base)
	if len(scan.Content) != 2 {
		t.Fatal("expected description added")
	}
}
