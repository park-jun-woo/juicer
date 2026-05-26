//ff:func feature=scan type=test control=sequence
//ff:what TestToOpenAPI_WithBaseNodeCov 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestToOpenAPI_WithBaseNodeCov(t *testing.T) {
	result := &ScanResult{}
	baseNode := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "paths"},
		{Kind: yaml.MappingNode},
	}}
	data, err := ToOpenAPI(result, baseNode)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("expected non-empty")
	}
}
