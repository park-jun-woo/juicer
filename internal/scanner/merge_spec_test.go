//ff:func feature=scan type=test control=sequence
//ff:what TestMergeSpec 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeSpec(t *testing.T) {
	scanNode := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "paths"},
		{Kind: yaml.MappingNode},
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
