//ff:func feature=scan type=test control=sequence
//ff:what TestFindMappingValue_Found 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestFindMappingValue_Found(t *testing.T) {
	val := &yaml.Node{Kind: yaml.ScalarNode, Value: "v"}
	node := &yaml.Node{
		Kind: yaml.MappingNode,
		Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "key"},
			val,
		},
	}
	result := findMappingValue(node, "key")
	if result != val {
		t.Fatal("expected val node")
	}
}

