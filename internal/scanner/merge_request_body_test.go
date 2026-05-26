//ff:func feature=scan type=test control=sequence
//ff:what TestMergeRequestBody_Both 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeRequestBody_Both(t *testing.T) {
	scan := &yaml.Node{Kind: yaml.MappingNode}
	base := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "description"},
		{Kind: yaml.ScalarNode, Value: "Body desc"},
	}}
	result := mergeRequestBody(scan, base)
	if result != scan {
		t.Fatal("expected scan")
	}
}
