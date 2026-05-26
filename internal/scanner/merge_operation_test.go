//ff:func feature=scan type=test control=sequence
//ff:what TestMergeOperation_BothPresent 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeOperation_BothPresent(t *testing.T) {
	scanOp := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "responses"},
		{Kind: yaml.MappingNode},
	}}
	baseOp := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "summary"},
		{Kind: yaml.ScalarNode, Value: "Get users"},
	}}
	result := mergeOperation(scanOp, baseOp)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

