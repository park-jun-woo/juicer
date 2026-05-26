//ff:func feature=scan type=test control=sequence
//ff:what TestMergeBaseOpsInto_BaseOnlyCov 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeBaseOpsInto_BaseOnlyCov(t *testing.T) {
	result := &yaml.Node{Kind: yaml.MappingNode}
	baseOps := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "post"},
		{Kind: yaml.MappingNode},
	}}
	scanOps := &yaml.Node{Kind: yaml.MappingNode}
	added := mergeBaseOpsInto(result, baseOps, scanOps)
	if !added["post"] {
		t.Fatal("expected post")
	}
}
