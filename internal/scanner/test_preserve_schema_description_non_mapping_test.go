//ff:func feature=scan type=test control=sequence
//ff:what TestPreserveSchemaDescription_NonMapping 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestPreserveSchemaDescription_NonMapping(t *testing.T) {
	preserveSchemaDescription(&yaml.Node{Kind: yaml.ScalarNode}, &yaml.Node{Kind: yaml.ScalarNode})
}
