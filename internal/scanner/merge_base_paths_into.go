//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what base paths 중 등록된 것만 mergedPaths에 추가한다
package scanner

import "gopkg.in/yaml.v3"

func mergeBasePathsInto(mergedPaths *yaml.Node, basePaths *yaml.Node, registered map[string]bool) {
	if basePaths == nil || basePaths.Kind != yaml.MappingNode {
		return
	}
	for i := 0; i+1 < len(basePaths.Content); i += 2 {
		path := basePaths.Content[i].Value
		ops := basePaths.Content[i+1]
		filteredOps, hasRegistered := filterOperations(ops, path, registered)
		if hasRegistered {
			mergedPaths.Content = append(mergedPaths.Content, basePaths.Content[i], filteredOps)
		}
	}
}
