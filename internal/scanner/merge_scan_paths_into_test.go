//ff:func feature=scan type=test control=sequence
//ff:what TestMergeScanPathsInto 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeScanPathsInto(t *testing.T) {
	merged := &yaml.Node{Kind: yaml.MappingNode}
	scan := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "/api/users"},
		{Kind: yaml.MappingNode},
	}}
	mergeScanPathsInto(merged, scan)
	if len(merged.Content) != 2 {
		t.Fatalf("expected 2 nodes, got %d", len(merged.Content))
	}
}
