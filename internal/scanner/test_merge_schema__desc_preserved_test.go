//ff:func feature=scan type=extract control=sequence
//ff:what TestMergeSchemas_DescPreserved 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMergeSchemas_DescPreserved(t *testing.T) {
	scanSchemas := &yaml.Node{Kind: yaml.MappingNode}
	scanSchema := &yaml.Node{Kind: yaml.MappingNode}
	scanSchema.Content = append(scanSchema.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Value: "type"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "object"},
	)
	scanSchemas.Content = append(scanSchemas.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Value: "User"},
		scanSchema,
	)

	baseSchemas := &yaml.Node{Kind: yaml.MappingNode}
	baseSchema := &yaml.Node{Kind: yaml.MappingNode}
	baseSchema.Content = append(baseSchema.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Value: "type"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "object"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "description"},
		&yaml.Node{Kind: yaml.ScalarNode, Value: "A user in the system"},
	)
	baseSchemas.Content = append(baseSchemas.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Value: "User"},
		baseSchema,
	)

	merged := mergeSchemas(scanSchemas, baseSchemas)

	userSchema := findMappingValue(merged, "User")
	if userSchema == nil {
		t.Fatal("expected User schema in merged result")
	}

	desc := findMappingValue(userSchema, "description")
	if desc == nil || desc.Value != "A user in the system" {
		t.Error("expected description preserved from base schema")
	}
}
