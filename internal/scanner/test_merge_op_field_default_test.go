//ff:func feature=scan type=test control=sequence
//ff:what TestMergeOpField_Default 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeOpField_Default(t *testing.T) {
	scanVal := &yaml.Node{Kind: yaml.ScalarNode, Value: "custom"}
	baseOp := &yaml.Node{Kind: yaml.MappingNode}
	result := mergeOpField("x-custom", scanVal, baseOp)
	if result != scanVal {
		t.Fatal("expected scan value for unknown key")
	}
}
