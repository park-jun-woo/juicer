//ff:func feature=scan type=test control=sequence
//ff:what TestPreserveParamDescription_NoBaseDesc 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestPreserveParamDescription_NoBaseDesc(t *testing.T) {
	scan := &yaml.Node{Kind: yaml.MappingNode}
	base := &yaml.Node{Kind: yaml.MappingNode}
	preserveParamDescription(scan, base)
}
