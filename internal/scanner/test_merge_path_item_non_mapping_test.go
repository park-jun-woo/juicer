//ff:func feature=scan type=test control=sequence
//ff:what TestMergePathItem_NonMapping 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergePathItem_NonMapping(t *testing.T) {
	scanOps := &yaml.Node{Kind: yaml.ScalarNode}
	baseOps := &yaml.Node{Kind: yaml.ScalarNode}
	result := mergePathItem(scanOps, baseOps)
	if result != scanOps {
		t.Fatal("expected scan returned")
	}
}
