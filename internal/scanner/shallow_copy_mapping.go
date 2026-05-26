//ff:func feature=scan type=extract control=sequence
//ff:what MappingNode의 최상위 키-값을 복사한다
package scanner

import "gopkg.in/yaml.v3"

func shallowCopyMapping(node *yaml.Node) *yaml.Node {
	if node == nil || node.Kind != yaml.MappingNode {
		return &yaml.Node{Kind: yaml.MappingNode}
	}
	cp := &yaml.Node{Kind: yaml.MappingNode}
	cp.Content = append(cp.Content, node.Content...)
	return cp
}
