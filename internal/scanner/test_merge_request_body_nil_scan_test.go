//ff:func feature=scan type=test control=sequence
//ff:what TestMergeRequestBody_NilScan 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeRequestBody_NilScan(t *testing.T) {
	base := &yaml.Node{Kind: yaml.MappingNode}
	if mergeRequestBody(nil, base) != base {
		t.Fatal("expected base")
	}
}
