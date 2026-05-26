//ff:func feature=scan type=test control=sequence
//ff:what TestMergeOpField_PreserveSummary 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeOpField_PreserveSummary(t *testing.T) {
	scanVal := &yaml.Node{Kind: yaml.ScalarNode, Value: "scan-summary"}
	baseOp := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "summary"},
		{Kind: yaml.ScalarNode, Value: "base-summary"},
	}}
	result := mergeOpField("summary", scanVal, baseOp)
	if result.Value != "base-summary" {
		t.Fatalf("expected base-summary, got %s", result.Value)
	}
}

