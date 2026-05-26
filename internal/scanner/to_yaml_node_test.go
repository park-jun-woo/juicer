//ff:func feature=scan type=test control=sequence
//ff:what TestToYAMLNode_Map 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestToYAMLNode_Map(t *testing.T) {
	node := toYAMLNode(map[string]any{"key": "value"})
	if node.Kind != yaml.MappingNode {
		t.Fatal("expected mapping node")
	}
}

