//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what scan responses를 기반으로 base에서 description을 보존한다
package scanner

import "gopkg.in/yaml.v3"

func mergeResponses(scanResps *yaml.Node, baseResps *yaml.Node) *yaml.Node {
	if scanResps == nil {
		return baseResps
	}
	if baseResps == nil || baseResps.Kind != yaml.MappingNode || scanResps.Kind != yaml.MappingNode {
		return scanResps
	}
	for i := 0; i+1 < len(scanResps.Content); i += 2 {
		statusCode := scanResps.Content[i].Value
		baseResp := findMappingValue(baseResps, statusCode)
		if baseResp == nil {
			continue
		}
		preserveDescription(scanResps.Content[i+1], baseResp)
	}
	return scanResps
}
