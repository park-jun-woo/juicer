//ff:func feature=scan type=test control=sequence
//ff:what TestSetMappingValue_Replace 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestSetMappingValue_Replace(t *testing.T) {
	node := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "key"},
		{Kind: yaml.ScalarNode, Value: "old"},
	}}
	newVal := &yaml.Node{Kind: yaml.ScalarNode, Value: "new"}
	setMappingValue(node, "key", newVal)
	if node.Content[1].Value != "new" {
		t.Fatal("expected replaced value")
	}
}
