//ff:func feature=scan type=test control=sequence
//ff:what TestAppendScanOnlyOps 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestAppendScanOnlyOps(t *testing.T) {
	result := &yaml.Node{Kind: yaml.MappingNode}
	scanOps := &yaml.Node{
		Kind: yaml.MappingNode,
		Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "get"},
			{Kind: yaml.ScalarNode, Value: "op1"},
			{Kind: yaml.ScalarNode, Value: "post"},
			{Kind: yaml.ScalarNode, Value: "op2"},
		},
	}
	added := map[string]bool{"get": true}
	appendScanOnlyOps(result, scanOps, added)
	if len(result.Content) != 2 {
		t.Fatalf("expected 2 nodes, got %d", len(result.Content))
	}
	if result.Content[0].Value != "post" {
		t.Fatalf("expected post, got %s", result.Content[0].Value)
	}
}
