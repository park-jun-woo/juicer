//ff:func feature=scan type=test control=sequence
//ff:what TestMergeOpField_Responses 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeOpField_Responses(t *testing.T) {
	scanVal := &yaml.Node{Kind: yaml.MappingNode}
	baseOp := &yaml.Node{Kind: yaml.MappingNode}
	result := mergeOpField("responses", scanVal, baseOp)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}
