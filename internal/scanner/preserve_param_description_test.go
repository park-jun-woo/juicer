//ff:func feature=scan type=test control=sequence
//ff:what TestPreserveParamDescription_NilBase 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestPreserveParamDescription_NilBase(t *testing.T) {
	scan := &yaml.Node{Kind: yaml.MappingNode}
	preserveParamDescription(scan, nil)
}
