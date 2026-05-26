//ff:func feature=scan type=test control=sequence
//ff:what TestFilterOperations 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestFilterOperations(t *testing.T) {
	ops := &yaml.Node{
		Kind: yaml.MappingNode,
		Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "get"},
			{Kind: yaml.ScalarNode, Value: "op1"},
			{Kind: yaml.ScalarNode, Value: "post"},
			{Kind: yaml.ScalarNode, Value: "op2"},
		},
	}
	reg := map[string]bool{"get\t/api/users": true}
	filtered, hasAny := filterOperations(ops, "/api/users", reg)
	if !hasAny {
		t.Fatal("expected hasAny")
	}
	if len(filtered.Content) != 2 {
		t.Fatalf("expected 2 nodes, got %d", len(filtered.Content))
	}
}
