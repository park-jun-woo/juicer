//ff:func feature=scan type=test control=sequence
//ff:what TestSetComponentSchemas_New 테스트
package scanner

import (
	"gopkg.in/yaml.v3"
	"testing"
)

func TestSetComponentSchemas_New(t *testing.T) {
	node := &yaml.Node{Kind: yaml.MappingNode}
	schemas := &yaml.Node{Kind: yaml.MappingNode}
	setComponentSchemas(node, schemas)
	comp := findMappingValue(node, "components")
	if comp == nil {
		t.Fatal("expected components")
	}
}
