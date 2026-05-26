//ff:func feature=scan type=test control=sequence
//ff:what TestMergeParameters_NilScan 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeParameters_NilScan(t *testing.T) {
	base := &yaml.Node{Kind: yaml.SequenceNode}
	result := mergeParameters(nil, base)
	if result != base {
		t.Fatal("expected base")
	}
}
