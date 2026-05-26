//ff:func feature=scan type=test control=sequence
//ff:what TestShallowCopyMapping 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestShallowCopyMapping(t *testing.T) {
	node := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "k"},
		{Kind: yaml.ScalarNode, Value: "v"},
	}}
	cp := shallowCopyMapping(node)
	if len(cp.Content) != 2 {
		t.Fatal("expected 2")
	}
}
