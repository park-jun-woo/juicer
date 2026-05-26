//ff:func feature=scan type=extract control=selection
//ff:what any 값을 yaml.Node로 재귀 변환한다 (맵 키 정렬 보장)
package scanner

import (
	"fmt"
	"os"
	"sort"

	"gopkg.in/yaml.v3"
)

func toYAMLNode(v any) *yaml.Node {
	switch val := v.(type) {
	case map[string]any:
		node := &yaml.Node{Kind: yaml.MappingNode}
		keys := make([]string, 0, len(val))
		for k := range val {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			keyNode := &yaml.Node{Kind: yaml.ScalarNode, Value: k}
			valNode := toYAMLNode(val[k])
			node.Content = append(node.Content, keyNode, valNode)
		}
		return node

	case []map[string]any:
		node := &yaml.Node{Kind: yaml.SequenceNode}
		for _, item := range val {
			node.Content = append(node.Content, toYAMLNode(item))
		}
		return node

	case []string:
		node := &yaml.Node{Kind: yaml.SequenceNode}
		for _, s := range val {
			node.Content = append(node.Content, &yaml.Node{Kind: yaml.ScalarNode, Value: s})
		}
		return node

	case *yaml.Node:
		return val

	case string:
		return &yaml.Node{Kind: yaml.ScalarNode, Value: val}

	case bool:
		n := &yaml.Node{Kind: yaml.ScalarNode}
		if val {
			n.Value = "true"
		} else {
			n.Value = "false"
		}
		n.Tag = "!!bool"
		return n

	case int:
		return &yaml.Node{Kind: yaml.ScalarNode, Value: fmt.Sprintf("%d", val), Tag: "!!int"}

	default:
		// 폴백: yaml.Marshal → Decode
		b, err := yaml.Marshal(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "toYAMLNode: marshal failed: %v\n", err)
			return &yaml.Node{Kind: yaml.ScalarNode, Value: fmt.Sprintf("<marshal-error: %v>", err)}
		}
		var n yaml.Node
		if err := yaml.Unmarshal(b, &n); err != nil {
			fmt.Fprintf(os.Stderr, "toYAMLNode: unmarshal failed: %v\n", err)
			return &yaml.Node{Kind: yaml.ScalarNode, Value: fmt.Sprintf("<unmarshal-error: %v>", err)}
		}
		if n.Kind == yaml.DocumentNode && len(n.Content) > 0 {
			return n.Content[0]
		}
		return &yaml.Node{Kind: yaml.ScalarNode, Value: ""}
	}
}

