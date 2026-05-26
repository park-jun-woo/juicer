//ff:func feature=scan type=extract control=sequence
//ff:what TestMergeOperation_DescPreserved 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMergeOperation_DescPreserved(t *testing.T) {
	scanOp := &yaml.Node{Kind: yaml.MappingNode}
	scanOp.Content = append(scanOp.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Value: "responses"},
		&yaml.Node{Kind: yaml.MappingNode},
	)

	baseOp := &yaml.Node{Kind: yaml.MappingNode}
	baseOp.Content = append(baseOp.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Value: "summary"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "List all users"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "description"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "Returns a paginated list of users"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "operationId"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "listUsers"},
	)

	merged := mergeOperation(scanOp, baseOp)

	summary := findMappingValue(merged, "summary")
	if summary == nil || summary.Value != "List all users" {
		t.Error("expected summary preserved from base")
	}

	desc := findMappingValue(merged, "description")
	if desc == nil || desc.Value != "Returns a paginated list of users" {
		t.Error("expected description preserved from base")
	}

	opID := findMappingValue(merged, "operationId")
	if opID == nil || opID.Value != "listUsers" {
		t.Error("expected operationId preserved from base")
	}
}
