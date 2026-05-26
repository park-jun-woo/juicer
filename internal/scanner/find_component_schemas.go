//ff:func feature=scan type=extract control=sequence
//ff:what components.schemas 노드를 반환한다
package scanner

import "gopkg.in/yaml.v3"

func findComponentSchemas(node *yaml.Node) *yaml.Node {
	comp := findMappingValue(node, "components")
	if comp == nil {
		return nil
	}
	return findMappingValue(comp, "schemas")
}
