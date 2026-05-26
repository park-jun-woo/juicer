//ff:func feature=scan type=extract control=selection
//ff:what operation 필드 종류에 따라 base 값을 보존하며 병합한다
package scanner

import "gopkg.in/yaml.v3"

var preserveFromBase = map[string]bool{
	"operationId": true,
	"summary":     true,
	"description": true,
	"tags":        true,
}

func mergeOpField(key string, scanVal *yaml.Node, baseOp *yaml.Node) *yaml.Node {
	switch {
	case preserveFromBase[key]:
		baseVal := findMappingValue(baseOp, key)
		if baseVal != nil {
			return baseVal
		}
		return scanVal
	case key == "parameters":
		return mergeParameters(scanVal, findMappingValue(baseOp, "parameters"))
	case key == "requestBody":
		return mergeRequestBody(scanVal, findMappingValue(baseOp, "requestBody"))
	case key == "responses":
		return mergeResponses(scanVal, findMappingValue(baseOp, "responses"))
	default:
		return scanVal
	}
}
