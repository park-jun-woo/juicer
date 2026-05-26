//ff:func feature=scan type=test control=sequence
//ff:what TestMergeSchemas_Both 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeSchemas_Both(t *testing.T) {
	scan := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "User"},
		{Kind: yaml.MappingNode},
	}}
	base := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "User"},
		{Kind: yaml.MappingNode, Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "description"},
			{Kind: yaml.ScalarNode, Value: "User model"},
		}},
	}}
	result := mergeSchemas(scan, base)
	if result != scan {
		t.Fatal("expected scan")
	}
}
