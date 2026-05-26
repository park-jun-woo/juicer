//ff:func feature=scan type=extract control=sequence
//ff:what base 노드의 description을 scan 노드에 보존한다
package scanner

import "gopkg.in/yaml.v3"

func preserveDescription(scanNode *yaml.Node, baseNode *yaml.Node) {
	if scanNode == nil || baseNode == nil {
		return
	}
	if scanNode.Kind != yaml.MappingNode || baseNode.Kind != yaml.MappingNode {
		return
	}
	baseDesc := findMappingValue(baseNode, "description")
	if baseDesc == nil {
		return
	}
	scanDesc := findMappingValue(scanNode, "description")
	if scanDesc == nil {
		scanNode.Content = append(scanNode.Content,
			&yaml.Node{Kind: yaml.ScalarNode, Value: "description"},
			baseDesc)
		return
	}
	setMappingValue(scanNode, "description", baseDesc)
}
