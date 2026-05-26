//ff:func feature=scan type=test control=sequence
//ff:what TestPreserveDescription_NonMapping 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestPreserveDescription_NonMapping(t *testing.T) {
	preserveDescription(&yaml.Node{Kind: yaml.ScalarNode}, &yaml.Node{Kind: yaml.ScalarNode})
}
