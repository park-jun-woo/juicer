//ff:func feature=scan type=extract control=sequence
//ff:what schema 노드에서 최상위 및 properties 내 description을 보존한다
package scanner

import "gopkg.in/yaml.v3"

func preserveSchemaDescription(scanSchema *yaml.Node, baseSchema *yaml.Node) {
	if scanSchema == nil || baseSchema == nil {
		return
	}
	if scanSchema.Kind != yaml.MappingNode || baseSchema.Kind != yaml.MappingNode {
		return
	}
	preserveDescription(scanSchema, baseSchema)
	preservePropertyDescriptions(scanSchema, baseSchema)
}
