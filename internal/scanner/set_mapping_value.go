//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what MappingNode에 key-value를 설정한다 (기존이면 교체, 없으면 추가)
package scanner

import "gopkg.in/yaml.v3"

func setMappingValue(node *yaml.Node, key string, value *yaml.Node) {
	if node == nil || node.Kind != yaml.MappingNode {
		return
	}
	for i := 0; i+1 < len(node.Content); i += 2 {
		if node.Content[i].Value == key {
			node.Content[i+1] = value
			return
		}
	}
	node.Content = append(node.Content,
		&yaml.Node{Kind: yaml.ScalarNode, Value: key}, value)
}
