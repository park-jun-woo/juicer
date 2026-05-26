//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what MappingNode에서 key의 인덱스를 반환한다
package scanner

import "gopkg.in/yaml.v3"

func findMappingIndex(node *yaml.Node, key string) int {
	if node == nil || node.Kind != yaml.MappingNode {
		return -1
	}
	for i := 0; i+1 < len(node.Content); i += 2 {
		if node.Content[i].Value == key {
			return i
		}
	}
	return -1
}
