//ff:func feature=scan type=test control=sequence
//ff:what TestMergeOpField_Parameters 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeOpField_Parameters(t *testing.T) {
	scanVal := &yaml.Node{Kind: yaml.SequenceNode}
	baseOp := &yaml.Node{Kind: yaml.MappingNode}
	result := mergeOpField("parameters", scanVal, baseOp)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}
