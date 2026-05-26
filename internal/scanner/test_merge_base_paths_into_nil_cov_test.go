//ff:func feature=scan type=test control=sequence
//ff:what TestMergeBasePathsInto_NilCov 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeBasePathsInto_NilCov(t *testing.T) {
	merged := &yaml.Node{Kind: yaml.MappingNode}
	mergeBasePathsInto(merged, nil, nil)
	if len(merged.Content) != 0 {
		t.Fatal("expected empty")
	}
}
