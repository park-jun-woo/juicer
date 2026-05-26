//ff:func feature=scan type=test control=sequence
//ff:what TestMergeResponses_NilScan 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeResponses_NilScan(t *testing.T) {
	base := &yaml.Node{Kind: yaml.MappingNode}
	if mergeResponses(nil, base) != base {
		t.Fatal("expected base")
	}
}
