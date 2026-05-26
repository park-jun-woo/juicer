//ff:func feature=scan type=test control=sequence
//ff:what TestMergeOpField_PreserveNoBase 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeOpField_PreserveNoBase(t *testing.T) {
	scanVal := &yaml.Node{Kind: yaml.ScalarNode, Value: "scan-desc"}
	baseOp := &yaml.Node{Kind: yaml.MappingNode}
	result := mergeOpField("description", scanVal, baseOp)
	if result != scanVal {
		t.Fatal("expected scan value when base is missing")
	}
}
