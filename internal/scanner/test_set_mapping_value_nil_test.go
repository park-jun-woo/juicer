//ff:func feature=scan type=test control=sequence
//ff:what TestSetMappingValue_Nil 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestSetMappingValue_Nil(t *testing.T) {
	setMappingValue(nil, "key", &yaml.Node{})
}
