//ff:func feature=scan type=test control=sequence
//ff:what TestMergeSchemas_NilScan 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeSchemas_NilScan(t *testing.T) {
	base := &yaml.Node{Kind: yaml.MappingNode}
	if mergeSchemas(nil, base) != base {
		t.Fatal("expected base")
	}
}
