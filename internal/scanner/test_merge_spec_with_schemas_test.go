//ff:func feature=scan type=test control=sequence
//ff:what TestMergeSpec_WithSchemas 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeSpec_WithSchemas(t *testing.T) {
	schemas := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "User"},
		{Kind: yaml.MappingNode},
	}}
	comp := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "schemas"},
		schemas,
	}}
	scanNode := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "paths"},
		{Kind: yaml.MappingNode},
		{Kind: yaml.ScalarNode, Value: "components"},
		comp,
	}}
	baseNode := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "paths"},
		{Kind: yaml.MappingNode},
	}}
	sr := &ScanResult{}
	result := mergeSpec(scanNode, baseNode, sr)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}
