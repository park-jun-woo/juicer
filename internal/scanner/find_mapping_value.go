//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what MappingNode에서 key에 해당하는 value 노드를 반환한다
package scanner

import "gopkg.in/yaml.v3"

func findMappingValue(node *yaml.Node, key string) *yaml.Node {
	if node == nil || node.Kind != yaml.MappingNode {
		return nil
	}
	for i := 0; i+1 < len(node.Content); i += 2 {
		if node.Content[i].Value == key {
			return node.Content[i+1]
		}
	}
	return nil
}
