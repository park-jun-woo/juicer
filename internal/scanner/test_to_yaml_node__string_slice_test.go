//ff:func feature=scan type=extract control=sequence
//ff:what TestToYAMLNode_StringSlice 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestToYAMLNode_StringSlice(t *testing.T) {
	node := toYAMLNode([]string{"a", "b"})
	if node.Kind != yaml.SequenceNode {
		t.Fatal("expected sequence node")
	}
}
