//ff:func feature=scan type=extract control=sequence
//ff:what scan requestBody를 기반으로 base에서 description을 보존한다
package scanner

import "gopkg.in/yaml.v3"

func mergeRequestBody(scanBody *yaml.Node, baseBody *yaml.Node) *yaml.Node {
	if scanBody == nil {
		return baseBody
	}
	if baseBody == nil {
		return scanBody
	}
	preserveDescription(scanBody, baseBody)
	return scanBody
}
