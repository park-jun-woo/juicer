//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestFilterOperations_NonHTTPKeys — non-HTTP 키가 보존되는지 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestFilterOperations_NonHTTPKeys(t *testing.T) {
	ops := &yaml.Node{
		Kind: yaml.MappingNode,
		Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "summary"},
			{Kind: yaml.ScalarNode, Value: "User operations"},
			{Kind: yaml.ScalarNode, Value: "description"},
			{Kind: yaml.ScalarNode, Value: "Detailed description"},
			{Kind: yaml.ScalarNode, Value: "parameters"},
			{Kind: yaml.MappingNode},
			{Kind: yaml.ScalarNode, Value: "get"},
			{Kind: yaml.ScalarNode, Value: "op1"},
			{Kind: yaml.ScalarNode, Value: "post"},
			{Kind: yaml.ScalarNode, Value: "op2"},
		},
	}
	reg := map[string]bool{"get\t/api/users": true}
	filtered, hasAny := filterOperations(ops, "/api/users", reg)
	if !hasAny {
		t.Fatal("expected hasAny true")
	}

	// summary, description, parameters (3 non-HTTP) + get (1 registered) = 4 pairs = 8 nodes
	// post is dead spec → dropped
	if len(filtered.Content) != 8 {
		t.Fatalf("expected 8 nodes (4 pairs), got %d", len(filtered.Content))
	}

	keys := map[string]bool{}
	for i := 0; i+1 < len(filtered.Content); i += 2 {
		keys[filtered.Content[i].Value] = true
	}
	for _, k := range []string{"summary", "description", "parameters", "get"} {
		if !keys[k] {
			t.Errorf("expected key %s to be preserved", k)
		}
	}
	if keys["post"] {
		t.Error("post should have been dropped as dead spec")
	}
}
