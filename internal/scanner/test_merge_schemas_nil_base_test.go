//ff:func feature=scan type=test control=sequence
//ff:what TestMergeSchemas_NilBase 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeSchemas_NilBase(t *testing.T) {
	scan := &yaml.Node{Kind: yaml.MappingNode}
	if mergeSchemas(scan, nil) != scan {
		t.Fatal("expected scan")
	}
}
