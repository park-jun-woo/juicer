//ff:func feature=scan type=extract control=sequence
//ff:what TestToYAMLNode_MapSlice 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestToYAMLNode_MapSlice(t *testing.T) {
	node := toYAMLNode([]map[string]any{{"a": "b"}})
	if node.Kind != yaml.SequenceNode {
		t.Fatal("expected sequence node")
	}
}
