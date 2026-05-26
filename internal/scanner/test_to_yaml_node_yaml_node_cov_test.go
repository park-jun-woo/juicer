//ff:func feature=scan type=test control=sequence
//ff:what TestToYAMLNode_YAMLNodeCov 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestToYAMLNode_YAMLNodeCov(t *testing.T) {
	orig := &yaml.Node{Kind: yaml.ScalarNode, Value: "x"}
	if toYAMLNode(orig) != orig {
		t.Fatal("expected same node")
	}
}
