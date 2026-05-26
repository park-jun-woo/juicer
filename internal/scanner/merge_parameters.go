//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what scan parameters를 기반으로 base에서 description을 보존한다
package scanner

import "gopkg.in/yaml.v3"

func mergeParameters(scanParams *yaml.Node, baseParams *yaml.Node) *yaml.Node {
	if scanParams == nil {
		return baseParams
	}
	if baseParams == nil || baseParams.Kind != yaml.SequenceNode {
		return scanParams
	}
	baseByName := indexParamsByName(baseParams)
	for _, param := range scanParams.Content {
		name := findMappingValue(param, "name")
		if name == nil {
			continue
		}
		preserveParamDescription(param, baseByName[name.Value])
	}
	return scanParams
}
