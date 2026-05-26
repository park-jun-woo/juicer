//ff:func feature=scan type=test control=sequence
//ff:what TestFilterOperations_OnlyNonHTTPKeys — HTTP 메서드 없이 non-HTTP 키만 있을 때 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestFilterOperations_OnlyNonHTTPKeys(t *testing.T) {
	ops := &yaml.Node{
		Kind: yaml.MappingNode,
		Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "summary"},
			{Kind: yaml.ScalarNode, Value: "Summary text"},
			{Kind: yaml.ScalarNode, Value: "servers"},
			{Kind: yaml.MappingNode},
		},
	}
	reg := map[string]bool{}
	filtered, hasAny := filterOperations(ops, "/api/test", reg)
	// no HTTP method found → hasAny should be false
	if hasAny {
		t.Fatal("expected hasAny false when only non-HTTP keys")
	}
	// non-HTTP keys should still be preserved
	if len(filtered.Content) != 4 {
		t.Fatalf("expected 4 nodes (2 pairs), got %d", len(filtered.Content))
	}
}
