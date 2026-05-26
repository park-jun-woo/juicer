//ff:func feature=scan type=extract control=sequence
//ff:what base parameter의 description을 scan parameter에 보존한다
package scanner

import "gopkg.in/yaml.v3"

func preserveParamDescription(scanParam *yaml.Node, baseParam *yaml.Node) {
	if baseParam == nil {
		return
	}
	baseDesc := findMappingValue(baseParam, "description")
	if baseDesc == nil {
		return
	}
	scanDesc := findMappingValue(scanParam, "description")
	if scanDesc == nil {
		scanParam.Content = append(scanParam.Content,
			&yaml.Node{Kind: yaml.ScalarNode, Value: "description"},
			baseDesc)
		return
	}
	setMappingValue(scanParam, "description", baseDesc)
}
