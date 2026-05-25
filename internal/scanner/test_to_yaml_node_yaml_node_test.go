//ff:func feature=scan type=extract control=sequence
//ff:what TestToYAMLNode_YAMLNode 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestToYAMLNode_YAMLNode(t *testing.T) {
	n := &yaml.Node{Kind: yaml.ScalarNode, Value: "test"}
	result := toYAMLNode(n)
	if result != n {
		t.Fatal("expected same node")
	}
}
