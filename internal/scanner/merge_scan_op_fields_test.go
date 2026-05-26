//ff:func feature=scan type=test control=sequence
//ff:what TestMergeScanOpFields 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMergeScanOpFields(t *testing.T) {
	result := &yaml.Node{Kind: yaml.MappingNode}
	scanOp := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "responses"},
		{Kind: yaml.MappingNode},
	}}
	baseOp := &yaml.Node{Kind: yaml.MappingNode}
	added := mergeScanOpFields(result, scanOp, baseOp)
	if !added["responses"] {
		t.Fatal("expected responses added")
	}
}
