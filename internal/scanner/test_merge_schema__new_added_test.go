//ff:func feature=scan type=extract control=sequence
//ff:what TestMergeSchemas_NewAdded 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMergeSchemas_NewAdded(t *testing.T) {
	// scan에 User와 Post가 있고, base에는 User만 있음
	scanSchemas := &yaml.Node{Kind: yaml.MappingNode}
	scanSchemas.Content = append(scanSchemas.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Value: "User"},
		&yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "type"},
			{Kind: yaml.ScalarNode, Value: "object"},
		}},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "Post"},
		&yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "type"},
			{Kind: yaml.ScalarNode, Value: "object"},
		}},
	)

	baseSchemas := &yaml.Node{Kind: yaml.MappingNode}
	baseSchemas.Content = append(baseSchemas.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Value: "User"},
		&yaml.Node{Kind: yaml.MappingNode, Content: []*yaml.Node{
			{Kind: yaml.ScalarNode, Value: "type"},
			{Kind: yaml.ScalarNode, Value: "object"},
			{Kind: yaml.ScalarNode, Value: "description"},
			{Kind: yaml.ScalarNode, Value: "A user"},
		}},
	)

	merged := mergeSchemas(scanSchemas, baseSchemas)

	// User는 있어야 하고
	if findMappingValue(merged, "User") == nil {
		t.Error("expected User schema")
	}
	// Post도 있어야 함 (scan에서 새로 추가)
	if findMappingValue(merged, "Post") == nil {
		t.Error("expected Post schema from scan")
	}
}
