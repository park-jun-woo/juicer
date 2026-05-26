//ff:func feature=scan type=test control=sequence
//ff:what TestPreservePropertyDescriptions_NilProps 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestPreservePropertyDescriptions_NilProps(t *testing.T) {
	scan := &yaml.Node{Kind: yaml.MappingNode}
	base := &yaml.Node{Kind: yaml.MappingNode}
	preservePropertyDescriptions(scan, base)
}
