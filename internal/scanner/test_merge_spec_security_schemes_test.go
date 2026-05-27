//ff:func feature=scan type=test control=sequence
//ff:what TestMergeSpec_SecuritySchemes 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMergeSpec_SecuritySchemes(t *testing.T) {
	// scan has securitySchemes, base doesn't
	scanComp := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "securitySchemes"},
		{Kind: yaml.MappingNode, Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "bearerAuth"},
			{Kind: yaml.MappingNode, Content: []*yaml.Node{
				{Kind: yaml.ScalarNode, Value: "type"},
				{Kind: yaml.ScalarNode, Value: "http"},
			}},
		}},
	}}
	scanNode := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "paths"},
		{Kind: yaml.MappingNode},
		{Kind: yaml.ScalarNode, Value: "components"},
		scanComp,
	}}
	baseNode := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "paths"},
		{Kind: yaml.MappingNode},
	}}
	sr := &ScanResult{}
	result := mergeSpec(scanNode, baseNode, sr)
	secSchemes := findComponentSecuritySchemes(result)
	if secSchemes == nil {
		t.Fatal("expected securitySchemes in merged result")
	}
	bearerAuth := findMappingValue(secSchemes, "bearerAuth")
	if bearerAuth == nil {
		t.Fatal("expected bearerAuth in merged securitySchemes")
	}
}
