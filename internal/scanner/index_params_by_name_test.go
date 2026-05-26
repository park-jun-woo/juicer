//ff:func feature=scan type=test control=sequence
//ff:what TestIndexParamsByName 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestIndexParamsByName(t *testing.T) {
	param1 := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "name"},
		{Kind: yaml.ScalarNode, Value: "id"},
	}}
	param2 := &yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
		{Kind: yaml.ScalarNode, Value: "name"},
		{Kind: yaml.ScalarNode, Value: "page"},
	}}
	params := &yaml.Node{Kind: yaml.SequenceNode, Content: []*yaml.Node{param1, param2}}
	result := indexParamsByName(params)
	if result["id"] != param1 {
		t.Fatal("expected param1 for 'id'")
	}
	if result["page"] != param2 {
		t.Fatal("expected param2 for 'page'")
	}
}
