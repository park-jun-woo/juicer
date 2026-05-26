//ff:func feature=scan type=test control=sequence
//ff:what TestMergeOperation_NilScan 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeOperation_NilScan(t *testing.T) {
	baseOp := &yaml.Node{Kind: yaml.MappingNode}
	result := mergeOperation(nil, baseOp)
	if result != baseOp {
		t.Fatal("expected base when scan is nil")
	}
}
