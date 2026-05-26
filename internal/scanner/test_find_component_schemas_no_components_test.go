//ff:func feature=scan type=test control=sequence
//ff:what TestFindComponentSchemas_NoComponents 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestFindComponentSchemas_NoComponents(t *testing.T) {
	root := &yaml.Node{Kind: yaml.MappingNode}
	result := findComponentSchemas(root)
	if result != nil {
		t.Fatal("expected nil")
	}
}
