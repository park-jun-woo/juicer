//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what paths를 경로 알파벳순으로 정렬한 yaml.Node로 반환한다
package scanner

import (
	"sort"

	"gopkg.in/yaml.v3"
)

func sortedPaths(paths map[string]map[string]any) *yaml.Node {
	keys := make([]string, 0, len(paths))
	for k := range paths {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	node := &yaml.Node{Kind: yaml.MappingNode}
	for _, k := range keys {
		keyNode := &yaml.Node{Kind: yaml.ScalarNode, Value: k}
		valNode := toYAMLNode(paths[k])
		node.Content = append(node.Content, keyNode, valNode)
	}
	return node
}

