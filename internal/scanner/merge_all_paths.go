//ff:func feature=scan type=extract control=sequence
//ff:what base paths를 필터링하고 scan paths를 병합한다
package scanner

import "gopkg.in/yaml.v3"

func mergeAllPaths(basePaths *yaml.Node, scanPaths *yaml.Node, registered map[string]bool) *yaml.Node {
	mergedPaths := &yaml.Node{Kind: yaml.MappingNode}
	mergeBasePathsInto(mergedPaths, basePaths, registered)
	mergeScanPathsInto(mergedPaths, scanPaths)
	return mergedPaths
}
