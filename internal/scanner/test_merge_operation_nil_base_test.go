//ff:func feature=scan type=test control=sequence
//ff:what TestMergeOperation_NilBase 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeOperation_NilBase(t *testing.T) {
	scanOp := &yaml.Node{Kind: yaml.MappingNode}
	result := mergeOperation(scanOp, nil)
	if result != scanOp {
		t.Fatal("expected scan when base is nil")
	}
}
