//ff:func feature=scan type=convert control=sequence
//ff:what TestToYAMLNode 테스트
package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestToYAMLNode(t *testing.T) {
	t.Run("map", func(t *testing.T) {
		node := toYAMLNode(map[string]any{"key": "value"})
		if node.Kind != 4 { // MappingNode
			t.Error("expected mapping node")
		}
	})

	t.Run("slice of maps", func(t *testing.T) {
		node := toYAMLNode([]map[string]any{{"a": "b"}})
		if node.Kind != 2 { // SequenceNode
			t.Error("expected sequence node")
		}
	})

	t.Run("slice of strings", func(t *testing.T) {
		node := toYAMLNode([]string{"a", "b"})
		if node.Kind != 2 { // SequenceNode
			t.Error("expected sequence node")
		}
	})

	t.Run("string", func(t *testing.T) {
		node := toYAMLNode("hello")
		if node.Value != "hello" {
			t.Errorf("expected 'hello', got %q", node.Value)
		}
	})

	t.Run("bool true", func(t *testing.T) {
		node := toYAMLNode(true)
		if node.Value != "true" {
			t.Errorf("expected 'true', got %q", node.Value)
		}
	})

	t.Run("bool false", func(t *testing.T) {
		node := toYAMLNode(false)
		if node.Value != "false" {
			t.Errorf("expected 'false', got %q", node.Value)
		}
	})

	t.Run("int", func(t *testing.T) {
		node := toYAMLNode(42)
		if node.Value != "42" {
			t.Errorf("expected '42', got %q", node.Value)
		}
	})

	t.Run("yaml.Node passthrough", func(t *testing.T) {
		input := &yaml.Node{Kind: yaml.ScalarNode, Value: "test"}
		node := toYAMLNode(input)
		if node != input {
			t.Error("expected same node to be returned")
		}
	})

	t.Run("fallback", func(t *testing.T) {
		node := toYAMLNode(3.14)
		if node == nil {
			t.Error("expected non-nil node")
		}
	})
}
