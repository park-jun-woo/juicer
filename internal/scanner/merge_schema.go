//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what components/schemas를 병합한다 (스캔이 진실, base에서 description 보존)
package scanner

import "gopkg.in/yaml.v3"

func mergeSchemas(scanSchemas *yaml.Node, baseSchemas *yaml.Node) *yaml.Node {
	if scanSchemas == nil {
		return baseSchemas
	}
	if baseSchemas == nil || baseSchemas.Kind != yaml.MappingNode {
		return scanSchemas
	}
	if scanSchemas.Kind != yaml.MappingNode {
		return scanSchemas
	}
	for i := 0; i+1 < len(scanSchemas.Content); i += 2 {
		schemaName := scanSchemas.Content[i].Value
		baseSchema := findMappingValue(baseSchemas, schemaName)
		if baseSchema == nil {
			continue
		}
		preserveSchemaDescription(scanSchemas.Content[i+1], baseSchema)
	}
	return scanSchemas
}
