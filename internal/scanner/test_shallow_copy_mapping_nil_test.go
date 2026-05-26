//ff:func feature=scan type=test control=sequence
//ff:what TestShallowCopyMapping_Nil 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestShallowCopyMapping_Nil(t *testing.T) {
	cp := shallowCopyMapping(nil)
	if cp.Kind != yaml.MappingNode {
		t.Fatal("expected mapping")
	}
}
