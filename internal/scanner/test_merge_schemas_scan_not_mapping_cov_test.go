//ff:func feature=scan type=test control=sequence
//ff:what TestMergeSchemas_ScanNotMappingCov 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeSchemas_ScanNotMappingCov(t *testing.T) {
	scan := &yaml.Node{Kind: yaml.ScalarNode}
	base := &yaml.Node{Kind: yaml.MappingNode}
	if mergeSchemas(scan, base) != scan {
		t.Fatal("expected scan")
	}
}
