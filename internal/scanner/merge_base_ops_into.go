//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what base operations를 result에 넣되, scan에도 있으면 mergeOperation으로 병합한다
package scanner

import "gopkg.in/yaml.v3"

func mergeBaseOpsInto(result *yaml.Node, baseOps *yaml.Node, scanOps *yaml.Node) map[string]bool {
	added := map[string]bool{}
	for i := 0; i+1 < len(baseOps.Content); i += 2 {
		method := baseOps.Content[i].Value
		scanOp := findMappingValue(scanOps, method)
		if scanOp != nil {
			result.Content = append(result.Content, baseOps.Content[i], mergeOperation(scanOp, baseOps.Content[i+1]))
		} else {
			result.Content = append(result.Content, baseOps.Content[i], baseOps.Content[i+1])
		}
		added[method] = true
	}
	return added
}
