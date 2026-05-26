//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what schema properties 내의 각 필드별 description을 base에서 보존한다
package scanner

import "gopkg.in/yaml.v3"

func preservePropertyDescriptions(scanSchema *yaml.Node, baseSchema *yaml.Node) {
	scanProps := findMappingValue(scanSchema, "properties")
	baseProps := findMappingValue(baseSchema, "properties")
	if scanProps == nil || baseProps == nil {
		return
	}
	if scanProps.Kind != yaml.MappingNode || baseProps.Kind != yaml.MappingNode {
		return
	}
	for i := 0; i+1 < len(scanProps.Content); i += 2 {
		propName := scanProps.Content[i].Value
		baseProp := findMappingValue(baseProps, propName)
		if baseProp == nil {
			continue
		}
		preserveDescription(scanProps.Content[i+1], baseProp)
	}
}
