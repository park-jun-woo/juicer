//ff:func feature=scan type=test control=sequence
//ff:what TestFindMappingValue_NotFound 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestFindMappingValue_NotFound(t *testing.T) {
	node := &yaml.Node{Kind: yaml.MappingNode}
	result := findMappingValue(node, "missing")
	if result != nil {
		t.Fatal("expected nil")
	}
}
