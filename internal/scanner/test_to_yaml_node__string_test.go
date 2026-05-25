//ff:func feature=scan type=extract control=sequence
//ff:what TestToYAMLNode_String 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestToYAMLNode_String(t *testing.T) {
	node := toYAMLNode("hello")
	if node.Kind != yaml.ScalarNode || node.Value != "hello" {
		t.Fatal("expected scalar hello")
	}
}
