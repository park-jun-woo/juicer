package scanner

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestToYAMLNode_Map(t *testing.T) {
	node := toYAMLNode(map[string]any{"key": "value"})
	if node.Kind != yaml.MappingNode {
		t.Fatal("expected mapping node")
	}
}

func TestToYAMLNode_StringSlice(t *testing.T) {
	node := toYAMLNode([]string{"a", "b"})
	if node.Kind != yaml.SequenceNode {
		t.Fatal("expected sequence node")
	}
}

func TestToYAMLNode_String(t *testing.T) {
	node := toYAMLNode("hello")
	if node.Kind != yaml.ScalarNode || node.Value != "hello" {
		t.Fatal("expected scalar hello")
	}
}

func TestToYAMLNode_Bool(t *testing.T) {
	node := toYAMLNode(true)
	if node.Value != "true" {
		t.Fatal("expected true")
	}
}

func TestToYAMLNode_Int(t *testing.T) {
	node := toYAMLNode(42)
	if node.Value != "42" {
		t.Fatal("expected 42")
	}
}

func TestToYAMLNode_MapSlice(t *testing.T) {
	node := toYAMLNode([]map[string]any{{"a": "b"}})
	if node.Kind != yaml.SequenceNode {
		t.Fatal("expected sequence node")
	}
}

func TestToYAMLNode_YAMLNode(t *testing.T) {
	n := &yaml.Node{Kind: yaml.ScalarNode, Value: "test"}
	result := toYAMLNode(n)
	if result != n {
		t.Fatal("expected same node")
	}
}

func TestToYAMLNode_Fallback(t *testing.T) {
	node := toYAMLNode(3.14)
	if node == nil {
		t.Fatal("expected non-nil")
	}
}
