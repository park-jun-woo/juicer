//ff:func feature=scan type=test control=sequence
//ff:what TestMergeScanPathsInto_ExistingPath 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeScanPathsInto_ExistingPath(t *testing.T) {
	merged := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "/api/users"},
		{Kind: yaml.MappingNode, Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "get"},
			{Kind: yaml.MappingNode},
		}},
	}}
	scan := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "/api/users"},
		{Kind: yaml.MappingNode, Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "post"},
			{Kind: yaml.MappingNode},
		}},
	}}
	mergeScanPathsInto(merged, scan)
}
