//ff:func feature=scan type=test control=sequence
//ff:what TestMergeParameters_NilBase 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeParameters_NilBase(t *testing.T) {
	scan := &yaml.Node{Kind: yaml.SequenceNode}
	result := mergeParameters(scan, nil)
	if result != scan {
		t.Fatal("expected scan")
	}
}
