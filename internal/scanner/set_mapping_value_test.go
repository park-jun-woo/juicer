//ff:func feature=scan type=test control=sequence
//ff:what TestSetMappingValue_Add 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestSetMappingValue_Add(t *testing.T) {
	node := &yaml.Node{Kind: yaml.MappingNode}
	val := &yaml.Node{Kind: yaml.ScalarNode, Value: "v"}
	setMappingValue(node, "key", val)
	if len(node.Content) != 2 {
		t.Fatal("expected 2")
	}
}

