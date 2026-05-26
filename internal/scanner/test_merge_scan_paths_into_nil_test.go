//ff:func feature=scan type=test control=sequence
//ff:what TestMergeScanPathsInto_Nil 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestMergeScanPathsInto_Nil(t *testing.T) {
	merged := &yaml.Node{Kind: yaml.MappingNode}
	mergeScanPathsInto(merged, nil)
}
