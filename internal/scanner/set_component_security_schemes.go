//ff:func feature=scan type=extract control=sequence
//ff:what merged 노드에 components.securitySchemes를 설정한다
package scanner

import "gopkg.in/yaml.v3"

func setComponentSecuritySchemes(node *yaml.Node, schemes *yaml.Node) {
	comp := findMappingValue(node, "components")
	if comp == nil {
		comp = &yaml.Node{Kind: yaml.MappingNode}
		setMappingValue(node, "components", comp)
	}
	setMappingValue(comp, "securitySchemes", schemes)
}
