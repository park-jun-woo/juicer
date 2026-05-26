//ff:func feature=scan type=test control=sequence
//ff:what TestToYAMLNode_MapSliceCov 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestToYAMLNode_MapSliceCov(t *testing.T) {
	node := toYAMLNode([]map[string]any{{"a": "b"}})
	if node.Kind != yaml.SequenceNode {
		t.Fatal("expected sequence")
	}
}
