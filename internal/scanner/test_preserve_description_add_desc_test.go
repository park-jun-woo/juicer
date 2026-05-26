//ff:func feature=scan type=test control=sequence
//ff:what TestPreserveDescription_AddDesc 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestPreserveDescription_AddDesc(t *testing.T) {
	scan := &yaml.Node{Kind: yaml.MappingNode}
	base := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "description"},
		{Kind: yaml.ScalarNode, Value: "Base desc"},
	}}
	preserveDescription(scan, base)
	if len(scan.Content) != 2 {
		t.Fatalf("expected 2 nodes, got %d", len(scan.Content))
	}
}
