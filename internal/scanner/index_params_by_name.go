//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what base parameters를 name으로 인덱싱한다
package scanner

import "gopkg.in/yaml.v3"

func indexParamsByName(params *yaml.Node) map[string]*yaml.Node {
	byName := map[string]*yaml.Node{}
	for _, param := range params.Content {
		name := findMappingValue(param, "name")
		if name != nil {
			byName[name.Value] = param
		}
	}
	return byName
}
