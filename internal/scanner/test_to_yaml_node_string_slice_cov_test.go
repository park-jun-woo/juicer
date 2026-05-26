//ff:func feature=scan type=test control=sequence
//ff:what TestToYAMLNode_StringSliceCov 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestToYAMLNode_StringSliceCov(t *testing.T) {
	node := toYAMLNode([]string{"a", "b"})
	if node.Kind != yaml.SequenceNode {
		t.Fatal("expected sequence")
	}
}
