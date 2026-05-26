//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what scan paths를 mergedPaths에 병합한다 (기존이면 operation 수준 병합)
package scanner

import "gopkg.in/yaml.v3"

func mergeScanPathsInto(mergedPaths *yaml.Node, scanPaths *yaml.Node) {
	if scanPaths == nil || scanPaths.Kind != yaml.MappingNode {
		return
	}
	for i := 0; i+1 < len(scanPaths.Content); i += 2 {
		path := scanPaths.Content[i].Value
		scanOps := scanPaths.Content[i+1]
		existingIdx := findMappingIndex(mergedPaths, path)
		if existingIdx < 0 {
			mergedPaths.Content = append(mergedPaths.Content, scanPaths.Content[i], scanOps)
			continue
		}
		baseOps := mergedPaths.Content[existingIdx+1]
		mergedPaths.Content[existingIdx+1] = mergePathItem(scanOps, baseOps)
	}
}
