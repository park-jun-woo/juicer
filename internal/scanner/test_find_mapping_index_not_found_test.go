//ff:func feature=scan type=test control=sequence
//ff:what TestFindMappingIndex_NotFound 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestFindMappingIndex_NotFound(t *testing.T) {
	node := &yaml.Node{Kind: yaml.MappingNode}
	if idx := findMappingIndex(node, "missing"); idx != -1 {
		t.Fatalf("expected -1, got %d", idx)
	}
}
