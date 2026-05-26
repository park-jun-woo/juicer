//ff:func feature=scan type=test control=sequence
//ff:what TestAppendBaseOnlyPreserved 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestAppendBaseOnlyPreserved(t *testing.T) {
	result := &yaml.Node{Kind: yaml.MappingNode}
	baseOp := &yaml.Node{
		Kind: yaml.MappingNode,
		Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "summary"},
			{Kind: yaml.ScalarNode, Value: "Get users"},
			{Kind: yaml.ScalarNode, Value: "operationId"},
			{Kind: yaml.ScalarNode, Value: "getUsers"},
			{Kind: yaml.ScalarNode, Value: "x-custom"},
			{Kind: yaml.ScalarNode, Value: "val"},
		},
	}
	added := map[string]bool{"summary": true}
	appendBaseOnlyPreserved(result, baseOp, added)
	// operationId should be appended (preserved + not added)
	// x-custom should not be appended (not in preserveFromBase)
	if len(result.Content) != 2 {
		t.Fatalf("expected 2 nodes (1 key+val pair), got %d", len(result.Content))
	}
	if result.Content[0].Value != "operationId" {
		t.Fatalf("expected operationId, got %s", result.Content[0].Value)
	}
}
