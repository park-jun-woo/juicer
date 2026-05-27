//ff:func feature=scan type=test control=sequence
//ff:what TestMergeSpec_BaseSecuritySchemesPreserved 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMergeSpec_BaseSecuritySchemesPreserved(t *testing.T) {
	// base has securitySchemes, scan doesn't
	baseComp := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "securitySchemes"},
		{Kind: yaml.MappingNode, Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "apiKey"},
			{Kind: yaml.MappingNode, Content: []*yaml.Node{
				{Kind: yaml.ScalarNode, Value: "type"},
				{Kind: yaml.ScalarNode, Value: "apiKey"},
			}},
		}},
	}}
	scanNode := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "paths"},
		{Kind: yaml.MappingNode},
	}}
	baseNode := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "paths"},
		{Kind: yaml.MappingNode},
		{Kind: yaml.ScalarNode, Value: "components"},
		baseComp,
	}}
	sr := &ScanResult{}
	result := mergeSpec(scanNode, baseNode, sr)
	secSchemes := findComponentSecuritySchemes(result)
	if secSchemes == nil {
		t.Fatal("expected base securitySchemes preserved in merged result")
	}
	apiKey := findMappingValue(secSchemes, "apiKey")
	if apiKey == nil {
		t.Fatal("expected apiKey in merged securitySchemes")
	}
}
